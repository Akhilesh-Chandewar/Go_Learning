package main

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	password := "password123"

	// SHA-256 and SHA-512 (original hashes)
	hash := sha256.Sum256([]byte(password))
	hash512 := sha512.Sum512([]byte(password))

	fmt.Println("sha256:", hash)
	fmt.Println("sha512:", hash512)
	fmt.Printf("SHA-256 hash hex: %x\n", hash)
	fmt.Printf("SHA-512 hash hex: %x\n", hash512)

	// Generate salt
	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}
	fmt.Println("Original Salt (raw bytes):", salt)
	fmt.Printf("Original Salt (hex): %x\n", salt)

	// Hash the password with salt
	signUpHash := hashPassword(password, salt)

	// Simulate storing salt and hash in database
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt (base64):", saltStr)
	fmt.Println("Hashed Password (base64):", signUpHash)

	// Hash password without salt (just to compare)
	hashOriginalPassword := sha256.Sum256([]byte(password))
	fmt.Println("Hash of password (no salt):", base64.StdEncoding.EncodeToString(hashOriginalPassword[:]))

	// ---- Simulate login verification ----

	// Decode stored salt
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Unable to decode salt:", err)
		return
	}

	// Hash the login password with decoded salt
	loginPassword := "password123" // change this to something else to simulate failure
	loginHash := hashPassword(loginPassword, decodedSalt)

	// Compare
	if signUpHash == loginHash {
		fmt.Println("✅ Password is correct. You are logged in.")
	} else {
		fmt.Println("❌ Login failed. Please check user credentials.")
	}
}

// Generate a random 16-byte salt
func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// Hash the password with salt using SHA-256
func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
