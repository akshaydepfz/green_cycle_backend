package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/main/internal/handler"
	"example.com/main/models"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		PostUserHandler(w, r)
	} else if r.Method == http.MethodGet {
		GetUserHandler(w, r)
	} else {
		fmt.Println("Method Invalid")
	}

}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var users models.User

	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	id, err := handler.PostUser(users.Name, &users.Orders)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	users.ID = id

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	users, err := handler.GetUser()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(users)
}
