package main

import (
	"fmt"
)

func main() {
	a := 'a'
	a += -a
	fmt.Println(a)
}
