package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	// Connect to  PostgreSQL database
	connStr := "user=postgres password=DPramesh@4 dbname=middleware sslmode=disable"

	// Open a connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the table
	createTableQuery := `CREATE TABLE IF NOT EXISTS control(id SERIAL PRIMARY KEY,name VARCHAR(100),Gender VARCHAR(100),email VARCHAR(100))`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table created successfully!")
}
