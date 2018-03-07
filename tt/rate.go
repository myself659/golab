package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
)

const (
	app     = "finance.rate"
	appkey  = "30616"
	sign    = "1d3a351596bad41a568d21b0d34b8fd9"
	baseurl = "http://api.k780.com/"
)

// ExchageRate ExchageRate
type ExchageRate struct {
	gorm.Model
	Status string `json:"status"`
	Scur   string `gorm:"type:varchar(20)" json:"scur"`
	Tcur   string `gorm:"type:varchar(20)" json:"tcur"`
	Ratenm string `gorm:"type:varchar(20)" json:"ratenm"`
	Rate   string `gorm:"type:varchar(20)" json:"rate"`
	Update string `gorm:"type:varchar(20)" json:"update"`
}

type rateResp struct {
	Success string      `json:"success"`
	Result  ExchageRate `json:"result"`
}

// FetchRate  FetchRate
func FetchRate(src string, dst string) {
	url := baseurl + "?app=" + app + "&scur=" + src + "&tcur=" + dst + "&appkey=" + appkey + "&sign=" + sign
	fmt.Println(url)
	//url = "http://api.k780.com/?app=finance.rate&scur=USD&tcur=CNY&appkey=10003&sign=b59bc3ef6191eb9f747dd4e83c99f2a4"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	var rate rateResp
	err = json.Unmarshal(b, &rate)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("rate:", rate)

}

func main() {
	FetchRate("USD", "CNY")
}

/*
http://api.k780.com/?app=finance.rate&scur=USD&tcur=CNY&appkey=30616&sign=1d3a351596bad41a568d21b0d34b8fd9
http://api.k780.com/?app=finance.rate&scur=USD&tcur=CNY&appkey=10003&sign=b59bc3ef6191eb9f747dd4e83c99f2a4
*/
