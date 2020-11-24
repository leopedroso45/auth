package handler

import (
	"auth/dto"
	"auth/model"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		name := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		age := r.FormValue("age")

		user := model.NewUser(name, email, password, age)

		//_, err := dto.CreateUser(user)
		sla := dto.CreateUser(user)
		//if err != nil {
		if sla == true {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", 201)
	}
}
