package handler

import (
	"MatchingApp/internal/usecase"
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

type Handler struct {
	service  *usecase.UseCase
	tpl      *template.Template
	consumer *sarama.Consumer
	producer *sarama.SyncProducer
}

func NewHandler(service *usecase.UseCase, tpl *template.Template, consumer *sarama.Consumer, producer *sarama.SyncProducer) *Handler {
	return &Handler{service: service, tpl: tpl, consumer: consumer, producer: producer}
}

func (h *Handler) Handle() http.Handler {
	r := gin.Default()
	r.SetHTMLTemplate(h.tpl)
	r.Static("/MatchingApp/stylesheet", "./templates/stylesheet")

	r.GET("/MatchingApp", h.userHandler)
	r.GET("/MatchingApp/", h.userHandler)
	r.GET("/MatchingApp/registrationUser", h.userHandler)
	r.GET("/MatchingApp/registrationUser/", h.userHandler)
	r.POST("/MatchingApp/createUser", h.userHandler)
	r.POST("/MatchingApp/createUser/", h.userHandler)
	r.GET("/MatchingApp/login", h.userHandler)
	r.GET("/MatchingApp/login/", h.userHandler)
	r.POST("/MatchingApp/loginUser", h.userHandler)
	r.POST("/MatchingApp/loginUser/", h.userHandler)
	r.GET("/MatchingApp/validate", h.RequireAuth, h.Validate)
	r.GET("/MatchingApp/validate/", h.RequireAuth, h.Validate)

	r.POST("/MatchingApp/createPlaylist", h.RequireAuth, h.playlistHandler)
	r.POST("/MatchingApp/createPlaylist/", h.RequireAuth, h.playlistHandler)
	r.GET("/MatchingApp/addPlaylist", h.RequireAuth, h.playlistHandler)
	r.GET("/MatchingApp/addPlaylist/", h.RequireAuth, h.playlistHandler)
	r.GET("/MatchingApp/match", h.RequireAuth, h.playlistHandler)
	r.GET("/MatchingApp/match/", h.RequireAuth, h.playlistHandler)

	return r
}
