package handler

import (
	"MatchingApp/internal/model"
	"MatchingApp/internal/passwordHelpers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

func (h *Handler) userHandler(c *gin.Context) {
	log.Printf("%s request on %s", c.Request.Method, c.Request.RequestURI)

	switch c.Request.Method {
	case http.MethodPost:
		if (regexp.MustCompile(`/MatchingApp/createUser*`)).MatchString(c.Request.URL.String()) {
			h.createUser(c)
		} else {
			h.loginUser(c)
		}

	case http.MethodDelete:
	case http.MethodPut:
	case http.MethodGet:
		if (regexp.MustCompile(`/MatchingApp/login*`)).MatchString(c.Request.URL.String()) {
			h.login(c)
		} else if (regexp.MustCompile(`/MatchingApp/registrationUser*`)).MatchString(c.Request.URL.String()) {
			h.registrationUser(c)
		} else {
			h.registrationUser(c)
		}
	default:
		c.AbortWithStatus(http.StatusMethodNotAllowed)
	}
}

func (h *Handler) registrationUser(c *gin.Context) {
	log.Println("*****registrationUser running*****")
	err := h.tpl.ExecuteTemplate(c.Writer, "registration.html", nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
func (h *Handler) createUser(c *gin.Context) {
	fmt.Println("*****registerAuthHandler running*****")
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")

	var user model.User
	user.Username = username
	user.Password = passwordHelpers.PasswordGenerator(password)
	user.Email = email

	h.service.UserUseCase.CreateUser(&user)

	log.Printf("createUser is completed")
	c.Status(http.StatusCreated)
	h.login(c)
}

func (h *Handler) login(c *gin.Context) {
	log.Println("*****login running*****")
	err := h.tpl.ExecuteTemplate(c.Writer, "login.html", nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
func (h *Handler) loginUser(c *gin.Context) {
	fmt.Println("*****loginUser running*****")
	username := c.PostForm("username")
	password := c.PostForm("password")

	user, err := h.service.FindUserByUsername(username)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	if user.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	// Compare sent in password with saved users password
	checker := passwordHelpers.CheckPass(user.Password, password)

	if checker != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// Respond
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.HTML(http.StatusOK, "loginSuccessfully.html", gin.H{})

}

func (h *Handler) Validate(c *gin.Context) {
	user, _ := c.Get("user")

	// user.(models.User).Email    -->   to access specific data

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
