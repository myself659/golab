package main

import "fmt"

func isEqual(target int) func(int) bool {
	return func(input int) bool {
		return input == target
	}
}

func main() {
	eq50 := isEqual(50)
	fmt.Println(eq50(50))
	fmt.Println(eq50(60))
}
