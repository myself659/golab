package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var r io.Reader

	r = strings.NewReader("123456")
	//r = io.LimitReader(r, 5)
	io.Copy(os.Stdout, r)

	fmt.Println("\nend")
}
