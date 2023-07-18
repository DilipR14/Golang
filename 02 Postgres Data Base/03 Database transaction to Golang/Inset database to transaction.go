package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID     int
	Name   string
	Gender string
	Email  string
}

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

	// Insert data into the database within the transaction
	users := []User{
		{ID: 1, Name: "dilip", Gender: "Male", Email: "jilip@example.com"},
		{ID: 2, Name: "Sunker", Gender: "Male", Email: "aun@example.com"},
	}

	for _, user := range users {
		_, err = tx.Exec("INSERT INTO selection(id, name,Gender, email) VALUES ($1, $2, $3, $4)", user.ID, user.Name, user.Gender, user.Email)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data inserted successfully!")
}
