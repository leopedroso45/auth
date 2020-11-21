package migration

import (
	"auth/db"
	"auth/model"
)

func AutoMigration() {
	db := db.Connect()
	defer db.Close()
	db.AutoMigrate(model.Music{}, model.User{})
}
