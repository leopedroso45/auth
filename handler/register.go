package handler

import (
	"auth/model"
	"net/http"
	"time"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		user := model.User{}
		user.Username = r.FormValue("username")
		user.Email = r.FormValue("email")
		user.Password = r.FormValue("password")
		user.Age = r.FormValue("age")
		user.CreatedAt = time.Now()
		user.Collection = []model.Music{}

		_, err := user.AddUser()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", 201)
	}
}
