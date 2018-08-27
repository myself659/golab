package main

import (
	"fmt"
	"net"
)

func main() {
	c, err := net.Dial("tcp", "101.132.105.131:8545")
	fmt.Println(c, err)

}
