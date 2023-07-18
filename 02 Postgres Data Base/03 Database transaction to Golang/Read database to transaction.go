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
	Gender *string
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

	// Read the database
	query := "SELECT id, name, email, gender FROM selection"
	rows, err := tx.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		var gender sql.NullString
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &gender)
		if err != nil {
			log.Fatal(err)
		}
		if gender.Valid {
			user.Gender = &gender.String
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Gender: %v\n", user.ID, user.Name, user.Email, user.Gender)
	}
}
