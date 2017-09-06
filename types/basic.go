package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"time"
)

func num2str() {
	// s := string(123) // wrong way
	s := strconv.Itoa(123)
	fmt.Println(s)
}

func int642byte() {
	ts := time.Now().Unix()
	fmt.Println(time.Now().String())
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, ts)
	fmt.Println(buf)
	fmt.Printf("%x", buf)
}

func int2byte(n int) {
	bs := []byte(string(n))
	fmt.Println()
	fmt.Printf("%x", bs)
	fmt.Println()
}

func main() {
	num2str()
	int642byte()
	int2byte(0x32)
}
