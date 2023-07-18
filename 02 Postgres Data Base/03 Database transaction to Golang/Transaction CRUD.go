package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Connect to the PostgreSQL database
	connStr := "user=postgres password=DPramesh@4 dbname=middleware sslmode=disable"

	// Open a connection to the PostgreSQL server
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Begin the transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	// Perform database operations within the transaction
	// Example: Insert a new record
	_, err = tx.Exec("INSERT INTO selection (name, email) VALUES ($1, $2)", "Karan", "karan78@example.com")
	if err != nil {
		log.Fatal(err)
	}

	// Example: Update a record
	_, err = tx.Exec("UPDATE selection SET email = $1 WHERE id = $2", "john.doe@example.com", 1)
	if err != nil {
		log.Fatal(err)
	}

	// Example: Delete a record
	_, err = tx.Exec("DELETE FROM selection WHERE id = $1", 2)
	if err != nil {
		log.Fatal(err)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaction committed successfully!")
}
