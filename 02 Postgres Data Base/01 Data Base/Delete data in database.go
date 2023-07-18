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
	// Connect the database
	connStr := "user=postgres password=DPramesh@4 dbname=middleware sslmode=disable"

	// Open a connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Delete data from table
	deleteDataQuery := `
		DELETE FROM selection
		WHERE id = $1
	`

	personID := 1
	result, err := db.Exec(deleteDataQuery, personID)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data deleted successfully. Rows affected: %d\n", rowsAffected)
}
