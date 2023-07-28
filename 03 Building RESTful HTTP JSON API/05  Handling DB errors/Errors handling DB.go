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

func insertUser(db *sql.DB, req *userDetailsRequest) error {
	insertQuery := "INSERT INTO id (1, Dilip) VALUES ($1, $2)"
	_, err := db.Exec(insertQuery, req.UserID, req.Username)
	if err != nil {
		return fmt.Errorf("Failed to insert user details: %v", err)
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

	if err := insertUser(db, &req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User details inserted successfully")
}

func main() {
	http.HandleFunc("/user-details", userDetailsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
