package main

import (
	"fmt"
	"time"
)

func main() {
	sch := make(chan []string)

	go func() {
		cs := make([]string, 0)
		cs = append(cs, "abc")
		cs = append(cs, "def")
		fmt.Println("func:", cs, len(cs))
		sch <- cs

	}()

	fmt.Println(<-sch)
	<-time.After(time.Second)
}
