package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	d := crypto.Keccak256Hash([]byte("baz(uint32,bool)"))
	r := d[0:4]
	fmt.Printf("%0x", r)
	fmt.Println()

	d = crypto.Keccak256Hash([]byte("C()"))
	r = d[0:4]
	fmt.Printf("%0x", r)
	fmt.Println()

}
