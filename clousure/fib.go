package main

import (
	"fmt"
)

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

/*
想一个更方便出来
非递归方式不需要使用closure

*/

func main() {
	f := fib()
	//fmt.Println(f())
	//fmt.Println(f())
	fmt.Println("-------")
	for i := 0; i < 3; i++ {
		fmt.Println(f())
	}
}
