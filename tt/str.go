package main

import (
	"fmt"
)

func main() {
	x := 12345
	sx := fmt.Sprintf("%06d", x)
	fmt.Println(sx)

	fmt.Println()
	a := []byte{48, 49}
	fmt.Printf("%s", string(a))

}
