package main

import (
	"flag"
	"fmt"
	"net"
	"syscall"
)

var ipport = "47.52.55.153:443"
var n = flag.Int("n", 10000000000, "conn")

func unlimit() {
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		fmt.Println("Error Getting Rlimit ", err)
	}
	fmt.Println(rLimit)
	rLimit.Max = 999999
	rLimit.Cur = 999999
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		fmt.Println("Error Setting Rlimit ", err)
	}
	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		fmt.Println("Error Getting Rlimit ", err)
	}
	fmt.Println("Rlimit Final", rLimit)
}

func main() {
	fmt.Println(*n)
	unlimit()

	for i := 0; i < *n; i++ {
		conn, err := net.Dial("tcp", ipport)
		if err != nil {
			fmt.Println(err, conn)
			conn.Write([]byte("test"))

		}
	}

}
