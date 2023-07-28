package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	
	var jwtSecret = []byte("your-secret-key")

	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = "dilip"
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() 

	// Generate the token string

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		fmt.Println("Error generating JWT token:", err)
		return
	}

	fmt.Println("JWT Token:", tokenString)

	// Verify the JWT token
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		fmt.Println("Error parsing JWT token:", err)
		return
	}

	if parsedToken.Valid {
		fmt.Println("JWT Token is valid")
		claims := parsedToken.Claims.(jwt.MapClaims)
		fmt.Println("Username:", claims["username"])
		fmt.Println("Role:", claims["role"])
	} else {
		fmt.Println("JWT Token is invalid")
	}
}
