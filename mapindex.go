package main

import (
	"fmt"
)

func main() {

	urls := make(map[string]int)
	urls["a"] = 100

	v, ok := urls["b"]
	fmt.Println(v, ", ", ok)
	v, ok = urls["a"]
	fmt.Println(v, ", ", ok)
	fmt.Println(urls)

}
