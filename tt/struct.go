package main

import (
	"fmt"
)

type Emd struct {
	Name string
}

type Com struct {
	Emd
	a int
}

func main() {
	a := 1234
	fmt.Printf("%6d", a)
	fmt.Printf("%06d", a)
	data := struct {
		title string
		Data  struct {
			content string
		}
	}{
		title: "hello",
		Data: struct {
			content string
		}{
			content: "world",
		},
	}
	fmt.Println(data)

}
