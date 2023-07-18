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
	Email  string
	Gender string
}

func main() {
	// Connect to the PostgreSQL database
	connStr := "user=postgres password=DPramesh@4 dbname=middleware sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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

	// Update data in the database
	userToUpdate := User{
		ID:     7,
		Name:   "Dojo",
		Email:  "doja.doe@gmail.com",
		Gender: "Male",
	}

	_, err = tx.Exec("UPDATE selection SET name = $1, email = $2, gender = $3 WHERE id = $4",
		userToUpdate.Name, userToUpdate.Email, userToUpdate.Gender, userToUpdate.ID)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data updated successfully!")
}
