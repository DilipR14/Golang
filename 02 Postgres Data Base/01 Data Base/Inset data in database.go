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
	// Connect toPostgreSQL database
	connStr := "user=postgres password=DPramesh@4 dbname=middleware sslmode=disable"

	// Open a connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	insertDataQuery := `
		INSERT INTO control (name, gender, email)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	persons := []Person{

		{1, "Dilip", "Male", "dilip@gmail.com"},
		{2, "Kumar", "Male", "kumar007@gmail.com"},
		{3, "Babu", "Male", "babu67@gmail.com"},
		{4, "Karthik", "Male", "karthik34@gmail.com"},
		{5, "Bhasker", "Male", "bhasker49@gmail.com"},
	}

	for _, p := range persons {
		var id int
		err = db.QueryRow(insertDataQuery, p.Name, p.Gender, p.Email).Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Data inserted successfully with ID: %d\n", id)
	}
}
