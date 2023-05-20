package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/DeepjyotiSarmah/go_server/models"
	// "github.com/dgrijalva/jwt-go"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"msg": "Hello form go Server"}

	responseData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
	// json.NewEncoder(w).Encode(response)
}

func GetMe(w http.ResponseWriter, r *http.Request) {

	userId := r.Context().Value("userId")

	var user *models.User

	for _, usr := range users {
		if usr.UserId == userId {
			user = &usr
			break
		}
	}

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{})
		return
	}

	response := map[string]interface{}{
		"userId": user.UserId,
		"name":   user.Name,
		"email":  user.Email,
	}

	responseData, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)

}
