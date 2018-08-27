package main

import (
	"fmt"
	"time"
)

func main() {

	//norecv := make(chan int)
	go func() {
		i := 0
		for {

			select {
			case <-time.After(1 * time.Second):
				{
					i++
					fmt.Println(i, "second pass")
				}
			default:
				{
					//norecv <- i
					//有这个default 就不会走进了1s超时定时器分支了
				}
			}
		}
	}()

	<-time.After(5 * time.Second)
}
