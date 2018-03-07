package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// CoinCap CoinCap
type CoinCap struct {
	ID                string `gorm:"type:varchar(20),not null" json:"id"`
	Name              string `gorm:"type:varchar(20)" json:"name"`
	Symbol            string `gorm:"type:varchar(20)" json:"symbol"`
	Rank              string `gorm:"type:varchar(20)" json:"rank"`
	PriceUSD          string `gorm:"type:varchar(20)" json:"price_usd"`
	PriceBTC          string `gorm:"type:varchar(20)" json:"price_btc"`
	HvolumeUSD        string `gorm:"type:varchar(20)" json:"24h_volume_usd"`
	MarketCapUSD      string `gorm:"type:varchar(20)" json:"market_cap_usd"`
	AvailableSupply   string `gorm:"type:varchar(20)" json:"available_supply"`
	TotalSupply       string `gorm:"type:varchar(20)" json:"total_supply"`
	MaxSupply         string `gorm:"type:varchar(20)" json:"max_supply"`
	PercentChangeHour string `gorm:"type:varchar(20)" json:"percent_change_1h"`
	PercentChangeDay  string `gorm:"type:varchar(20)" json:"percent_change_24h"`
	PercentChangeWeek string `gorm:"type:varchar(20)" json:"percent_change_7d"`
	LastUpdated       string `gorm:"type:varchar(20)" json:"last_updated"`
	PriceCNY          string `gorm:"type:varchar(20)" json:"price_cny"`
	HvolumeCNY        string `gorm:"type:varchar(20)" json:"24h_volume_cny"`
	MarketCapCNY      string `gorm:"type:varchar(20)" json:"market_cap_cny"`
}

const (
	baseurl = "https://api.coinmarketcap.com/v1/ticker/"
)

func CoinCapByCoin(coin string, convert string) {
	url := baseurl + coin + "/?convert=" + convert
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
	var items []CoinCap
	fmt.Println(string(b))
	err = json.Unmarshal(b, &items)
	fmt.Println(err, items)
}

func CoinCaps(start int, limit int, convert string) {

	url := baseurl + "?start=" + strconv.Itoa(start) + "&limit=" + strconv.Itoa(limit) + "&convert=" + convert
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
	var items []CoinCap
	fmt.Println(string(b))
	err = json.Unmarshal(b, &items)
	fmt.Println(err, items)
}

func main() {
	CoinCapByCoin("bitcoin", "CNY")
	CoinCaps(0, 10, "CNY")
}
