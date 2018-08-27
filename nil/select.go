package main

import (
	"fmt"
	"time"
)

func sel(c1, c2 chan int) {

	for {
		select {
		case x, ok := <-c1:
			if !ok {
				c1 = nil
				fmt.Println("c1  empty")
			} else {
				fmt.Println("rec from c1:", x)
			}

		case x, ok := <-c2:
			if !ok {
				c2 = nil
				fmt.Println("c2  empty")
			} else {
				fmt.Println("rec from c2:", x)
			}
		default:
			fmt.Println("---default---")

		}
		if c1 == nil && c2 == nil {
			return
		}
	}
}

func main() {
	c1 := make(chan int, 2)
	c2 := make(chan int, 2)
	go sel(c1, c2)
	c1 <- 1
	c2 <- 2
	c1 <- 3
	c1 <- 5
	close(c1)

	time.Sleep(3 * time.Second)

}
