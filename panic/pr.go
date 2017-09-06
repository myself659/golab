package main

import (
	"fmt"
)

func panicdo(i int) {
	if i == 0 {
		defer panic("zero")
	}
	fmt.Println("test")
	//fmt.Println(50 / i)
	defer recover()
}

func main() {
	panicdo(0)
	r := recover()
	r = r
	fmt.Println(r)
}
