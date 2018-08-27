package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%T\n", math.Pi)

	const n = 500000000

	fmt.Printf("%T\n", n)
	a := 46
	fmt.Printf("%s", string(a))

}
