package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DeepjyotiSarmah/go_server/models"
)

var users = []models.User{}
var userCountId = 1

func SignUp(w http.ResponseWriter, r *http.Request) {

	var requestSignup struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestSignup); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, usr := range users {
		if usr.Email == requestSignup.Email {
			http.Error(w, "Email already exits", http.StatusForbidden)
			return
		}
	}

	user := models.User{
		Email:    requestSignup.Email,
		Password: requestSignup.Password,
		UserId:   strconv.Itoa(userCountId),
	}

	users = append(users, user)
	userCountId++

	response := map[string]string{"msg": "Success"}

	responseData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}
