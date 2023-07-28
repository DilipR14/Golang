package main

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/o1egl/paseto/v2"
)

func main() {
	
	fmt.Println("== JWT Token ==")

	var jwtSecret = []byte("your-secret-key")

	// Create a new JWT token
	jwtToken := jwt.New(jwt.SigningMethodHS256)

	jwtClaims := jwtToken.Claims.(jwt.MapClaims)
	jwtClaims["username"] = "dilip"
	jwtClaims["role"] = "admin"
	jwtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix() 

	// Generate the JWT token string
	jwtTokenString, err := jwtToken.SignedString(jwtSecret)
	if err != nil {
		fmt.Println("Error generating JWT token:", err)
		return
	}

	fmt.Println("JWT Token:", jwtTokenString)

	// Verify the JWT token
	parsedJWTToken, err := jwt.Parse(jwtTokenString, func(token *jwt.Token) (interface{}, error) {
		
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		fmt.Println("Error parsing JWT token:", err)
		return
	}

	if parsedJWTToken.Valid {
		fmt.Println("JWT Token is valid")
		claims := parsedJWTToken.Claims.(jwt.MapClaims)
		fmt.Println("Username:", claims["username"])
		fmt.Println("Role:", claims["role"])
	} else {
		fmt.Println("JWT Token is invalid")
	}

	fmt.Println()

	// PASETO Token
	fmt.Println("== PASETO Token ==")

	pasetoSecret := make([]byte, 32)
	_, err = rand.Read(pasetoSecret)
	if err != nil {
		fmt.Println("Error generating random key:", err)
		return
	}

	// Create a new PASETO token
	now := time.Now()
	expiration := now.Add(time.Hour * 24)
	pasetoToken := paseto.JSONToken{
		Expiration: expiration,
	}

	pasetoTokenString, err := paseto.NewV2().Encrypt(pasetoSecret, &pasetoToken, nil)
	if err != nil {
		fmt.Println("Error generating PASETO token:", err)
		return
	}

	fmt.Println("PASETO Token:", pasetoTokenString)

	var parsedPasetoToken paseto.JSONToken
	err = paseto.NewV2().Decrypt(pasetoTokenString, pasetoSecret, &parsedPasetoToken, nil)
	if err != nil {
		fmt.Println("Error parsing PASETO token:", err)
		return
	}

	// Verify the PASETO token
	err = paseto.NewV2().Verify(pasetoTokenString, pasetoSecret, &parsedPasetoToken, nil)
	if err != nil {
		fmt.Println("Error verifying PASETO token:", err)
		return
	}

	fmt.Println("PASETO Token is valid")
	fmt.Println("Expiration:", parsedPasetoToken.Expiration)
}
