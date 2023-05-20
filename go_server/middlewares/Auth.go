package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

const JWT_SECRET = "secret"

type contextKey string

const (
	userIDContextKey contextKey = "userId"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the Authorization header from the request
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]interface{}{"msg": "Missing auth header"})
			return
		}

		// Extract the JWT token from the Authorization header
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Verify the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verify the signing method and return the secret key
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(JWT_SECRET), nil
		})

		// Handle token parsing/validation errors
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]interface{}{"msg": "Incorrect token"})
			return
		}

		// Extract the user ID from the token claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if id, ok := claims["id"].(string); ok {
				// Set the user ID in the request context
				ctx := context.WithValue(r.Context(), userIDContextKey, id)
				r = r.WithContext(ctx)
				next(w, r)
				return
			}
		}

		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]interface{}{"msg": "Incorrect token"})
	}
}
