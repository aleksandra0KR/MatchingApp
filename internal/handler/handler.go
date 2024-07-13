package handler

import (
	"MatchingApp/internal/usecase"
	"net/http"
)

type Handler struct {
	service *usecase.UseCase
}

func NewHandler(service *usecase.UseCase) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Handle() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/MatchingApp", http.HandlerFunc(h.userHandler))
	mux.Handle("/MatchingApp/", http.HandlerFunc(h.userHandler))

	return mux
}
