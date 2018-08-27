package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, 14)
}
func CheckPasswordHash(hash, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}
func Example() {
	myPassword := []byte("password")
	hashed, err := HashPassword(myPassword)
	if err != nil {
		return
	}
	fmt.Println(string(hashed))
}

func main() {
	Example()
}
