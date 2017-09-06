package main

import "fmt"

func main() { // callnig level 0

	defer fmt.Println("program will crash, for panic 3 is stll active")

	defer func() { // calling level 1

		defer func() { // calling level 2

			fmt.Println(recover()) // 6
		}()

		// the level of panic 3 is 0.
		// the level of panic 6 is 1.
		defer fmt.Println("now, there are two active panics: 3 and 6")

		defer panic(6) // will suppress panic 5
		defer panic(5) // will suppress panic 4
		panic(4)       // will not suppress panic 3, for they have differrent levels
		// the level of panic 3 is 0.
		// the level of panic 4 is 1.
	}()

	defer fmt.Println("now, only panic 3 is active")

	defer panic(3) // will suppress panic 2
	defer panic(2) // will suppress panic 1
	panic(1)
}
