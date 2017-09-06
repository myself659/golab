package main

import (
	"fmt"
)

type stEA struct {
	b int
}
type stA struct {
	a    int
	data stEA
}

type stB struct {
	a int
	c int
}
type stC struct {
	ca int
	cc int
}

func main() {
	a := stA{}
	a.a = 1
	c := stC(a)
	fmt.Println(c)
	//b := stB(a)
	//fmt.Println(b)
}
