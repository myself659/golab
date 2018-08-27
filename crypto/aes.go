package main

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"
)

func main() {
	buf := []byte("The quick brown fox jumps over the lazy dog")
	// Generate highly secure random AES key
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	aesCipher, _ := aes.NewCipher(key)
	// Encrypt in-place.
	aesCipher.Encrypt(buf, buf)
	fmt.Printf("%s\n", buf)
}
