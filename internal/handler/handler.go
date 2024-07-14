package handler

import (
	"MatchingApp/internal/usecase"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

type Handler struct {
	service *usecase.UseCase
	tpl     *template.Template
}

func NewHandler(service *usecase.UseCase, tpl *template.Template) *Handler {
	return &Handler{service: service, tpl: tpl}
}

func (h *Handler) Handle() http.Handler {
	r := gin.Default()

	r.GET("/MatchingApp", h.userHandler)
	r.GET("/MatchingApp/", h.userHandler)
	r.POST("/MatchingApp/createUser", h.userHandler)
	r.POST("/MatchingApp/createUser/", h.userHandler)
	r.GET("/MatchingApp/login", h.userHandler)
	r.GET("/MatchingApp/login/", h.userHandler)
	r.POST("/MatchingApp/loginUser", h.userHandler)
	r.POST("/MatchingApp/loginUser/", h.userHandler)
	r.GET("/MatchingApp/validate", h.RequireAuth, h.Validate)
	r.GET("/MatchingApp/validate/", h.RequireAuth, h.Validate)

	return r
}
