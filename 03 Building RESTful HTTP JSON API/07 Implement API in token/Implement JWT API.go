
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var UserCredentials = map[string]string{
	"username1": "password1",
	"username2": "password2",

}

var jwtSecret = []byte("your-secret-key")

func GenerateJWTToken(username string) (string, error) {
	
	expiration := time.Now().Add(1 * time.Hour)

	claims := jwt.MapClaims{
		"username": username,
		"exp":      expiration.Unix(),
	}

	// Create the JWT token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	/
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	password, found := UserCredentials[user.Username]
	if !found || password != user.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := GenerateJWTToken(user.Username)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"access_token": token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/login", LoginHandler)

	
	fmt.Println("Server is running on https://github.com/DilipR14/golang")
	http.ListenAndServe(":8080", nil)
}
