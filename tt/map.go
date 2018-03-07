package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string)
	m["a"] = "aa"
	fmt.Println(m["a"])
	fmt.Println(m["bb"])
	v, ok := m["bb"]
	fmt.Println(v, ok)
}
