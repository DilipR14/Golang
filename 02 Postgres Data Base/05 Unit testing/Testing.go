package main

import (
	"database/sql"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	_ "github.com/lib/pq"
)

// Test CRUD operations
func TestCRUD(t *testing.T) {
	// Open a connection to the test database
	db, err := sql.Open("postgres", "dbname=middleware user=postgres password=DPramesh@4 host=localhost port=5432 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a new user
	user := User{
		ID:        rand.Intn(1000),
		FirstName: "Jos",
		LastName:  "D",
		Email:     "jo@gmail.com",
	}

	err = CreateUser(db, user)
	assert.NoError(t, err, "Failed to create user")

	// Retrieve the created user
	createdUser, err := GetUser(db, user.ID)
	assert.NoError(t, err, "Failed to get user")
	assert.Equal(t, user, createdUser, "Retrieved user is not equal to created user")

	// Update the user
	updatedUser := User{
		ID:        user.ID,
		FirstName: "Jag",
		LastName:  "S",
		Email:     "jan@gmail.com",
	}

	err = UpdateUser(db, updatedUser)
	assert.NoError(t, err, "Failed to update user")

	// Retrieve the updated user
	retrievedUser, err := GetUser(db, user.ID)
	assert.NoError(t, err, "Failed to get user")
	assert.Equal(t, updatedUser, retrievedUser, "Retrieved user is not equal to updated user")

	// Delete the user
	err = DeleteUser(db, user.ID)
	assert.NoError(t, err, "Failed to delete user")

	// Attempt to retrieve the deleted user
	deletedUser, err := GetUser(db, user.ID)
	assert.Error(t, err, "Expected error when getting deleted user")
	assert.EqualError(t, err, "User not found", "Unexpected error message when getting deleted user")
	assert.Equal(t, User{}, deletedUser, "Retrieved deleted user is not empty")
}

// User represents a user in the database
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

// CreateUser creates a new user in the database
func CreateUser(db *sql.DB, user User) error {
	query := "INSERT INTO persion (id, first_name, last_name, email) VALUES ($1, $2, $3, $4)"
	_, err := db.Exec(query, user.ID, user.FirstName, user.LastName, user.Email)
	return err
}

// GetUser retrieves a user from the database by ID
func GetUser(db *sql.DB, userID int) (User, error) {
	query := "SELECT selection , first_name, last_name, email FROM users WHERE id = $2"
	row := db.QueryRow(query, userID)

	var user User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("User not found")
		}
		return User{}, err
	}

	return user, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Run the tests
	err := testing.Main(func(_, _ string) (bool, error) {
		return true, nil
	}, []testing.InternalTest{{"TestCRUD", TestCRUD}})
	if err != nil {
		log.Fatal(err)
	}
}
