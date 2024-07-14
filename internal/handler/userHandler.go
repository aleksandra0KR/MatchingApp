package handler

import (
	"MatchingApp/internal/model"
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
	case http.MethodGet:
		h.registrationUser(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (h *Handler) registrationUser(w http.ResponseWriter, r *http.Request) {
	log.Println("*****registrationUser running*****")
	err := h.tpl.ExecuteTemplate(w, "registration.html", nil)
	if err != nil {
		return
	}
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("*****registerAuthHandler running*****")
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	var user model.User
	user.Username = username
	user.Password = password
	user.Email = email

	h.service.UserUseCase.CreateUser(&user)

	log.Printf("createUser is completed")
	w.WriteHeader(http.StatusCreated)
	_, err := fmt.Fprint(w, "user is created ")
	if err != nil {
		return
	}
}
