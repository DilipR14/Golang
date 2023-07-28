package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

type userDetailsRequest struct {
	UserID   int    `json:"user_id" validate:"required,min=1"`
	Username string `json:"username" validate:"required"`
}

type userDetailsResponse struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
}

func validateUserDetailsRequest(req *userDetailsRequest) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("Invalid value for '%s' field", err.Field())
		}
	}
	return nil
}

func userDetailsHandler(w http.ResponseWriter, r *http.Request) {
	
	var req userDetailsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := validateUserDetailsRequest(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Open a connection to the database
	db, err := sql.Open("postgres", "dbname=middleware user=postgres password=DPramesh@4 host=localhost port=5432 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := `
		SELECT user_id, username
		FROM users
		WHERE user_id = $1
	`
	row := db.QueryRow(query, req.UserID)

	// Create a userDetailsResponse 
	var user userDetailsResponse
	err = row.Scan(
		&user.UserID,
		&user.Username,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to fetch user details", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", "dbname=middleware user=postgres password=DPramesh@4 host=localhost port=5432 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the users table 
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS ID (
			user_id SERIAL PRIMARY KEY,
			username VARCHAR(255) UNIQUE NOT NULL
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/user-details", userDetailsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
