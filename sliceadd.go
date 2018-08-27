package main

import (
	"fmt"
)

func main() {
	sa := []int{1, 2, 3}
	sb := []int{2, 4, 6}
	//sa = sa + sb
	for _, i := range sb {
		sa = append(sa, i)
	}
	fmt.Println(sa)
	for _, i := range sa {
		fmt.Println(i)

	}
}

/*
Building Resilient Services in Go
*/

