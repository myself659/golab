package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 12345
	sx := fmt.Sprintf("%06d", x)
	fmt.Println(sx)

	fmt.Println()
	a := []byte{48, 49}
	fmt.Printf("%s\n", string(a))
	n := 32
	fmt.Printf("%s\n", strconv.FormatUint(uint64(n), 10))
	fmt.Printf("%s\n", strconv.FormatUint(uint64(n), 16))

}
