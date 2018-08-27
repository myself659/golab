package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		i := 0
		j := 0
		for {
			select {
			case <-time.After(1 * time.Second):
				{
					fmt.Println("i=", i)
				}
			default:
				{
					select {
					case <-time.After(2 * time.Second):
						{
							fmt.Println("2s pass ")
						}
					case ch <- j:
						{
							j++
						}
					}
				}
			}
		}
	}()

	<-time.After(10 * time.Second)
}
