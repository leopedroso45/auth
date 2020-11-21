package model

import (
	database "auth/db"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model

	ID         uint      `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Username   string    `gorm:"type:varchar(50);not null;unique"json:"username,omitempty"`
	Email      string    `gorm:"size:50;not null;unique"json:"email,omitempty"`
	Password   string    `gorm:"size:255;not null;"`
	Age        string    `gorm:"size:2;not null;"json:"age,omitempty"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	Collection []Music   `json:"collection,omitempty"`
}

func (user *User) AddUser() (*User, error) {
	conn := database.Connect()
	defer conn.Close()

	var err error
	err = conn.Debug().Model(&User{}).Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}
