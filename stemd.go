package main

import (
	"encoding/json"
	"fmt"
)

type emd struct {
	Name string
}

type con struct {
	emd
	Pwd string
}

func main() {
	var c con
	c.Name = "name"
	b, err := json.Marshal(c)
	fmt.Println(string(b))
	var cj con
	err = json.Unmarshal(b, &cj)
	fmt.Println(cj)
	fmt.Println(err)

	fmt.Printf("%v", c)
	//fmt.Printf("%+V", c)
}
