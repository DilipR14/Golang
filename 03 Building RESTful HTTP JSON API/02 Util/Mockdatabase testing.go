package main

import (
	"fmt"
	"log"
	"net/http"
)

type MockDatabase struct {
	DbName   string
	Password string
	Port     int
	Host     string
	Data     map[string]interface{}
}

func (mdb *MockDatabase) Create(data interface{}) error {
	
	return nil
}

func (mdb *MockDatabase) Read(id string) (interface{}, error) {
	
	return nil, nil
}

func (mdb *MockDatabase) Update(id string, data interface{}) error {
	
	return nil
}

func (mdb *MockDatabase) Delete(id string) error {

	return nil
}

func APIHandler(db *MockDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		fmt.Fprintf(w, "Mock Database: %s\n", db.DbName)
		fmt.Fprintf(w, "Password: %s\n", db.Password)
		fmt.Fprintf(w, "Port: %d\n", db.Port)
		fmt.Fprintf(w, "Host: %s\n", db.Host)
	}
}

func main() {
	// Create a new instance of the MockDatabase
	mockDB := &MockDatabase{
		DbName:   "middleware",
		Password: "DPramesh@4",
		Port:     5432,
		Host:     "localhost",
		Data:     make(map[string]interface{}),
	}

	http.HandleFunc("/api", APIHandler(mockDB))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
