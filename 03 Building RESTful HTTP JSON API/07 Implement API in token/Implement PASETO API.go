package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/o1egl/paseto"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var UserCredentials = map[string]string{
	"username1": "password1",
	"username2": "password2",
}

func GeneratePasetoToken(username string) (string, error) {
	
	secretKey := []byte("my-secret-key")

	// Create a new PASETO builder.
	builder := paseto.NewV2()

	payload := map[string]interface{}{
		"sub": "golang",
		"exp": "dilip",
	}

	encodedToken, err := builder.Encrypt(secretKey, payload, nil)
	if err != nil {
		return "", err
	}

	return encodedToken, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

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

	token, err := GeneratePasetoToken(user.Username)
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
