package main

import (
	handle "auth/handler"
	"log"
	"net/http"
)

func main() {
	//Handlers to implement
	http.HandleFunc("/signing", handle.Signing)
	http.HandleFunc("/welcome", handle.Welcome)
	http.HandleFunc("/refresh", handle.Refresh)

	//start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}
