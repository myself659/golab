package main

import (
	"fmt"
)

type ssFunc func(s string) string

var ffs []ssFunc

func regFunc(f ssFunc) {
	ffs = append(ffs, f)
}

func add1Func(start string) ssFunc {
	return func(s string) string {
		return s + "-1" + start
	}
}

func add2Func(suffix string) ssFunc {
	return func(s string) string {
		return s + "-2" + suffix
	}
}

func main() {

	fmt.Println("start")
	regFunc(add2Func(".txt"))
	regFunc(add1Func(".pdf"))
	genies := "0"
	for _, f := range ffs {
		temp := f(genies)
		fmt.Println(temp)
	}

}

/* 一层套一层调用 更是巧妙  */
