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

// Define a struct for the transfer request parameters
type transferRequest struct {
	From   int     `json:"from" validate:"required,min=1"`
	To     int     `json:"to" validate:"required,min=1"`
	Amount float64 `json:"amount" validate:"required,gt=0"`
}

// Define a custom validator for validating transferRequest struct
func validateTransferRequest(req *transferRequest) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("Invalid value for '%s' field", err.Field())
		}
	}
	return nil
}

func transferMoneyHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body into transferRequest struct
	var req transferRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate request parameters
	if err := validateTransferRequest(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", "dbname=middleware user=postgres password=DPramesh@4 host=localhost port=5432 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Perform the money transfer
	// ... (Write your transfer logic here)

	// Return success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Money transfer successful")
}

func main() {
	http.HandleFunc("/transfer", transferMoneyHandler)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
