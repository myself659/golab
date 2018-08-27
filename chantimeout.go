package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		i := 0
		j := 0
		for {
			select {
			case <-time.After(1 * time.Second):
				{
					i++
					fmt.Println("i=", i)
				}
			case <-time.After(2 * time.Second): // never  get time out
				{
					j++
					fmt.Println("j=", j)
				}
			}
		}
	}()
	<-time.After(10 * time.Second)
}
