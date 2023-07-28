package main

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/o1egl/paseto/v2"
)

func main() {
	
	secretKey := make([]byte, 32)
	_, err := rand.Read(secretKey)
	if err != nil {
		fmt.Println("Error generating random key:", err)
		return
	}

	// Create a new PASETO token
	now := time.Now()
	expiration := now.Add(time.Hour * 24)
	token := paseto.JSONToken{
		Expiration: expiration,
	}

	pasetoToken, err := paseto.NewV2().Encrypt(secretKey, token, nil)
	if err != nil {
		fmt.Println("Error generating PASETO token:", err)
		return
	}

	fmt.Println("PASETO Token:", pasetoToken)

	
	var parsedPasetoToken paseto.JSONToken
	err = paseto.NewV2().Decrypt(pasetoToken, secretKey, &parsedPasetoToken, nil)
	if err != nil {
		fmt.Println("Error parsing PASETO token:", err)
		return
	}

	fmt.Println("PASETO Token is valid")
	fmt.Println("Expiration:", parsedPasetoToken.Expiration)
}
