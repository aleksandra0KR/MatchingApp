package handler

import (
	"MatchingApp/internal/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (h *Handler) userHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s request on %s", r.Method, r.RequestURI)

	switch r.Method {
	case http.MethodPost:
		h.createUser(w, r)
	case http.MethodDelete:
	case http.MethodPut:
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	h.service.UserUseCase.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	log.Printf("createUser is completed")
	w.WriteHeader(http.StatusCreated)
	_, err = fmt.Fprint(w, "user is created ")
	if err != nil {
		return
	}
}
