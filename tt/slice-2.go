package main

import (
	"fmt"
)

func a() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)  // commonTags := labelsToTags(app.Labels)
	y := append(x, 2) // Tags: append(commonTags, labelsToTags(d.Labels)...)
	z := append(x, 3) // Tags: append(commonTags, labelsToTags(d.Labels)...)
	fmt.Println(y, z)
}

func deepcopy(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

func b() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2) // commonTags := labelsToTags(app.Labels)
	y := deepcopy(x)
	y = append(y, 3)  // Tags: append(commonTags, labelsToTags(d.Labels)...)
	z := append(x, 4) // Tags: append(commonTags, labelsToTags(d.Labels)...)
	fmt.Println(y, z)
}

func main() {
	a()
	b()
}
