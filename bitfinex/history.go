package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type tradeItem [6]float64

func main() {
	url := "https://api.bitfinex.com/v2/candles/trade:1m:tBTCUSD/hist?limit=10"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var items []tradeItem
	err = json.Unmarshal(b, &items)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(items)

}
