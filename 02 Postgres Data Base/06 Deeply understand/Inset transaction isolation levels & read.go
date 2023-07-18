package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = "DPramesh@4"
	dbName     = "middleware"
)

var db *sql.DB

func main() {
	// Connect to the database
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a table for testing
	createTableQuery := `CREATE TABLE IF NOT EXISTS Kumar (id SERIAL PRIMARY KEY, name VARCHAR(100), balance INT)`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	setIsolationLevel("READ COMMITTED")

	go transactionOne()
	go transactionTwo()

	var input string
	fmt.Println("Press Enter to exit")
	fmt.Scanln(&input)
}

func setIsolationLevel(level string) {
	query := fmt.Sprintf("SET TRANSACTION ISOLATION LEVEL %s", level)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func transactionOne() {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Insert new data
	_, err = tx.Exec("INSERT INTO Kumar (name, balance) VALUES ($1, $2)", "Dilip", 500)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func transactionTwo() {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	var balance int
	err = tx.QueryRow("SELECT balance FROM Kumar WHERE id = 1").Scan(&balance)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaction Two - Balance:", balance)

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
