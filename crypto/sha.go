package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

func Hash(tag string, data []byte) []byte {
	h := hmac.New(sha512.New512_256, []byte(tag))
	h.Write(data)
	return h.Sum(nil)
}

func ExampleHash() error {
	tag := "hashing file for storage key"
	contents, err := ioutil.ReadFile("sha.txt")
	if err != nil {
		return err
	}
	digest := Hash(tag, contents)
	fmt.Println(hex.EncodeToString(digest))
	return nil
}

func main() {
	ExampleHash()
	fileDigest := Hash("fileNode", []byte("hello, world"))
	metaDigest := Hash("metadataNode", []byte("hello, world"))
	fmt.Printf("%x\n%x\n", fileDigest, metaDigest)
}
