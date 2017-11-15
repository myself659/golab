package main

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

func main() {
	input := "abcdefghijkl"
	scanner := bufio.NewScanner(strings.NewReader(input))
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		fmt.Printf("split: %t\t%d\t%s\n", atEOF, len(data), data)
		if atEOF {
			return 0, nil, errors.New("bad luck")
		}
		return 0, nil, nil
	}
	scanner.Split(split)
	buf := make([]byte, 12)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)
	for scanner.Scan() {
		fmt.Printf("text:%s\n", scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Printf("error: %s\n", scanner.Err())
	}
}
