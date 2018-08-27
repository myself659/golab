package main

import (
	"fmt"
)

func print(s string) {
	fmt.Println(s)
}
func getFunc() func(string) {
	return print
}

func main() {
	// 返回函数直接进行调用
	getFunc()("hello world")
	f1 := getFunc()
	f1("f1")
	f2 := getFunc()
	f2("f2")
	fmt.Println(f1, f2)
}
