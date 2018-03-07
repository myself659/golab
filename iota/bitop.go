package main

import "fmt"

type Month int

const (
	// 1 << 0 ==> 1
	January  Month = 1 << iota
	February       // 1 << 1 ==> 2
	March          // 1 << 2 ==> 4
	April          // 1 << 3 ==> 8
	May            // 1 << 4 ==> 16
	June           // ...
	July
	August
	September
	October
	November
	December
	// Break the iota chain here.
	// AllMonths will have only
	// the assigned month values,
	// not the iota's.
	AllMonths = January | February |
		March | April | May | June |
		July | August | September |
		October | November |
		December
)

func main() {
	fmt.Println(February, AllMonths)
}
