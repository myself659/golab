package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func main() {
	in := []byte("hello world")
	out := sha1.Sum(in)
	fmt.Printf("%x", out)
	fmt.Println()
	outslice := out[0:len(out)]
	outstr := hex.EncodeToString(outslice)
	fmt.Println(outstr)
}
