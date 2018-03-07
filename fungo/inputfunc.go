package main

import (
	"fmt"
)

func filter(i int) bool {
	if i%2 == 0 {
		return true
	}
	return false
}

type filterFunc func(i int) bool

func output(ai []int, f filterFunc) {
	for _, i := range ai {
		if f(i) {
			fmt.Println(i)
		}
	}
}

func main() {
	fmt.Println("start")
	ai := []int{1, 2, 4, 7, 9, 8}
	output(ai, filter)
}
