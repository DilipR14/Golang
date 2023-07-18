package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Person struct {
	ID     int
	Name   string
	Gender string
	Email  string
}

func main() {
	// Connect to database
	connStr := "user=postgres password=DPramesh@4 dbname=middleware sslmode=disable"

	// Open a connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Update data in the table
	updateDataQuery := `
		UPDATE selection
		SET name = $1, gender = $2, email = $3
		WHERE id = $4
	`

	person := Person{
		ID:     5,
		Name:   "Deva",
		Gender: "Male",
		Email:  "dev67@gmail.com",
	}

	result, err := db.Exec(updateDataQuery, person.Name, person.Gender, person.Email, person.ID)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data updated successfully. Rows affected: %d\n", rowsAffected)
}
