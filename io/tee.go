package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {

	r := strings.NewReader("some io.Reader stream to be read\n")
	var buf bytes.Buffer
	tee := io.TeeReader(r, &buf)
	fmt.Println(tee)
	fmt.Println(&buf)

}
