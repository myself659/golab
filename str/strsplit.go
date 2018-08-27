package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aa, bb, c"
	sad := s + "abc" +
		"cd" +
		"ef"
	fmt.Println(sad)
	as := strings.Split(s, ",")
	fmt.Println(len(as))
	if len(as) == 3 {
		fmt.Println(as[0])
	}
	fmt.Printf("%T", as[0])
}
