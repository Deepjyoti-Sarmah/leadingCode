package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/DeepjyotiSarmah/go_server/models"
	"github.com/dgrijalva/jwt-go"
)

const JWT_SECRET = "secret"

func Login(w http.ResponseWriter, r *http.Request) {

	var requestLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestLogin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user *models.User

	for _, usr := range users {
		if usr.Email == requestLogin.Email {
			user = &usr
			break
		}
	}

	if user == nil {
		http.Error(w, "User not found", http.StatusForbidden)
		return
	}

	if user.Password != requestLogin.Password {
		http.Error(w, "Incorrect Password", http.StatusForbidden)
		return
	}

	//Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.UserId,
	})

	signedToken, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"token": signedToken}

	responseData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)

}
