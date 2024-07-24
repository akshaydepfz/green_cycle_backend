package main

import (
	"net/http"

	"example.com/main/database"
	"example.com/main/internal/repository"
)

func main() {
	database.Initdb()
	http.HandleFunc("/users", repository.UserHandler)
	http.ListenAndServe(":8080", nil)
}
