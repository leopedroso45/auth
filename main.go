package main

import (
	handle "auth/handler"
	"auth/model/migration"
	"log"
	"net/http"
)

func main() {
	// Run Migration
	migration.AutoMigration()

	//Handlers to implement
	http.HandleFunc("/register", handle.Register)
	http.HandleFunc("/signing", handle.Signing)
	http.HandleFunc("/welcome", handle.Welcome)
	http.HandleFunc("/refresh", handle.Refresh)

	//start the server on port 8000
	log.Fatal(http.ListenAndServe(":8080", nil))
}
