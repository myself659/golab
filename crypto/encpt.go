package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

func Encrypt(data []byte, key [32]byte) ([]byte, error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, err
	}
	return gcm.Seal(nonce, nonce, data, nil), nil
}

func Decrypt(ciphertext []byte, key [32]byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
}

func main() {
	data := []byte("123456")
	var key [32]byte
	key[0] = 11
	key[2] = 22
	cbyte, err := Encrypt(data, key)
	if err != nil {
		fmt.Println("Encrypt:", cbyte)
		return
	}
	//key[1] = 11
	mbyte, err := Decrypt(cbyte, key)
	if err != nil {
		fmt.Println("Decrypt:", mbyte)
		return
	}
	fmt.Println(cbyte)
	fmt.Println(mbyte)

}
