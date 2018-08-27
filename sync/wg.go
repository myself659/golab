package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func do(i int) {
	wg.Add(1)
	defer wg.Done()
	fmt.Println("i:", i)
}

func main() {
	for i := 0; i < 10; i++ {
		go do(i)
	}

	<-time.After(6 * time.Second)
	wg.Wait()
}
