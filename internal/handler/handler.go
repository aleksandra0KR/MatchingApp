package handler

import (
	"MatchingApp/internal/usecase"
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
	mux := http.NewServeMux()
	mux.Handle("/MatchingApp", http.HandlerFunc(h.userHandler))
	mux.Handle("/MatchingApp/", http.HandlerFunc(h.userHandler))
	mux.Handle("/MatchingApp/createUser", http.HandlerFunc(h.userHandler))
	mux.Handle("/MatchingApp/createUser/", http.HandlerFunc(h.userHandler))

	return mux
}
