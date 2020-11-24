package model

import (
	"encoding/json"
	"time"
)

type User struct {
	ID         string   `json:"id,omitempty"`
	Username   string `json:"username,omitempty"`
	Email      string `json:"email,omitempty"`
	Password   string
	Age        string    `json:"age,omitempty"`
	CreatedAt  time.Time `son:"created_at,omitempty"`
	Collection []Music   `json:"collection,omitempty"`
}
func NewUser(username, email, password, age string) User {
	return User{
		Username:   username,
		Email:      email,
		Password:   password,
		Age:        age,
		CreatedAt:  time.Now(),
	}
}


func (user *User) toJson() ([]byte, error) {
	userJson, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	return userJson, nil
}

func (user *User) setPassword(newPass string) string {
	user.Password = newPass
	return user.Password
}