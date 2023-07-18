package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=DPramesh@4 dbname=middleware sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = resetTables(db)
	if err != nil {
		log.Fatal(err)
	}

	err = resetSequences(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("API reset complete.")
}

func resetTables(db *sql.DB) error {
	_, err := db.Exec("TRUNCATE TABLE selection, control, Kumar CASCADE")
	if err != nil {
		return err
	}
	return nil
}

func resetSequences(db *sql.DB) error {
	tables := []string{"selection", "control", "Kumar"}

	for _, table := range tables {
		query := fmt.Sprintf("SELECT setval(pg_get_serial_sequence('%s', 'id'), coalesce(max(id), 0) + 1, false) FROM %s", table, table)
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}
