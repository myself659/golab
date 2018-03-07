package main
import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	fmt.Println(len(m))
	m["a"] = 0
	fmt.Println(len(m))
	m["b"] = 1
	fmt.Println(len(m))
}
