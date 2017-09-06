package main

import (
	"fmt"
	"os"
)

func a() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)  // commonTags := labelsToTags(app.Labels)
	y := append(x, 2) // Tags: append(commonTags, labelsToTags(d.Labels)...)
	z := append(x, 3) // Tags: append(commonTags, labelsToTags(d.Labels)...)
	fmt.Println(y, z)
}

func b() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2)  // commonTags := labelsToTags(app.Labels)
	y := append(x, 3) // Tags: append(commonTags, labelsToTags(d.Labels)...)
	z := append(x, 4) // Tags: append(commonTags, labelsToTags(d.Labels)...)
	fmt.Println(y, z)
}

func init() {
	fmt.Println("init add ")
}

func main() {
	fmt.Println(os.Args)
	a()
	b()
}
