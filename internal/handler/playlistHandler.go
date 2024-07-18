package handler

import (
	"MatchingApp/internal/model"
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
		} else {
			//	h.(c)
		}

	case http.MethodDelete:
	case http.MethodPut:
	case http.MethodGet:
		if (regexp.MustCompile(`/MatchingApp/login*`)).MatchString(c.Request.URL.String()) {
			//		h.login(c)
		} else if (regexp.MustCompile(`/MatchingApp/addPlaylist*`)).MatchString(c.Request.URL.String()) {
			h.addPlaylist(c)
		} else {
			//	h.Validate(c)
		}
	default:
		c.AbortWithStatus(http.StatusMethodNotAllowed)
	}
}

func (h *Handler) addPlaylist(c *gin.Context) {
	log.Println("*****addPlaylist running*****")
	err := h.tpl.ExecuteTemplate(c.Writer, "enterPlaylist.html", nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
func (h *Handler) createPlaylist(c *gin.Context) {
	fmt.Println("*****registerAuthHandler running*****")
	user, _ := c.Get("user")
	playlistId := c.PostForm("playlistId")
	username := c.PostForm("username")
	var playlist model.Playlist
	playlist.UserName = username
	playlist.UserID = user.(model.User).ID
	playlist.PlaylistKey = playlistId

	h.service.PlaylistUseCase.CreatePlaylist(&playlist)

	log.Printf("createPlaylist is completed")
	c.Status(http.StatusCreated)
	_, err := c.Writer.Write([]byte("playlist is created "))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
