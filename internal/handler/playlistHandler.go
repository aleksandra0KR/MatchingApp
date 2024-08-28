package handler

import (
	"MatchingApp/internal/model"
	kafka2 "MatchingApp/internal/model/kafka"
	"MatchingApp/kafka"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
)

func (h *Handler) playlistHandler(c *gin.Context) {
	log.Printf("%s request on %s", c.Request.Method, c.Request.RequestURI)

	switch c.Request.Method {
	case http.MethodPost:
		if (regexp.MustCompile(`/MatchingApp/createPlaylist*`)).MatchString(c.Request.URL.String()) {
			h.createPlaylist(c)
		}
	case http.MethodDelete:
	case http.MethodPut:
	case http.MethodGet:
		if (regexp.MustCompile(`/MatchingApp/match`)).MatchString(c.Request.URL.String()) {
			h.matchingPlaylist(c)
		} else if (regexp.MustCompile(`/MatchingApp/addPlaylist*`)).MatchString(c.Request.URL.String()) {
			h.addPlaylist(c)
		}
	default:
		c.AbortWithStatus(http.StatusMethodNotAllowed)
	}
}

// addPlaylist godoc
// @Summary redirect to post request with html page
// @Description Render the template to add a playlist.
// @Tags playlists
// @Accept  json
// @Produce  html
// @Router /addPlaylist [get]
func (h *Handler) addPlaylist(c *gin.Context) {
	log.Println("*****addPlaylist running*****")
	err := h.tpl.ExecuteTemplate(c.Writer, "enterPlaylist.html", nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

// createPlaylist godoc
// @Summary Add a new playlist
// @Description Create a new playlist for the authenticated user in their profile based on which program will find match.
// @Tags playlists
// @Accept  json
// @Produce  json
// @Param playlistID formData string true "Playlist ID"
// @Param username formData string true "Username"
// @Success 200 {object} string "Playlist created successfully"
// @Failure 500 {object} string "Internal Server Error"
// @Router /createPlaylist [post]
func (h *Handler) createPlaylist(c *gin.Context) {
	fmt.Println("*****registerAuthHandler running*****")
	user, _ := c.Get("user")
	playlistId := c.PostForm("playlistID")
	username := c.PostForm("username")
	var playlist model.Playlist
	playlist.UserName = username
	playlist.UserID = user.(model.User).ID
	playlist.PlaylistKey = playlistId
	kafka.SendMessage(h.producer, "", kafka2.Message{username, user.(model.User).ID.String(), playlistId})
	// h.service.PlaylistUseCase.CreatePlaylist(&playlist)

	log.Printf("createPlaylist is completed")
	/*c.Status(http.StatusCreated)
	_, err := c.Writer.Write([]byte("playlist is created "))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}*/
	message := kafka.ReadMessage(h.consumer, "")
	log.Println("*****matchingPlaylist running*****")
	err := h.tpl.ExecuteTemplate(c.Writer, "match.html", message)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	//c.HTML(http.StatusOK, "loginSuccessfully.html", gin.H{})
}

// matchingPlaylist godoc
// @Summary Match a playlist
// @Description Match a playlist and render the template.
// @Tags playlists
// @Accept  json
// @Produce  html
// @Router /match [get]
func (h *Handler) matchingPlaylist(c *gin.Context) {
	fmt.Println("ewf")
	message := kafka.ReadMessage(h.consumer, "")
	fmt.Println(message)
	log.Println("*****matchingPlaylist running*****")
	err := h.tpl.ExecuteTemplate(c.Writer, "match.html", message)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
