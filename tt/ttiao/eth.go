package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"time"
)

var url = "https://etherscan.io/txs?a=0x48ee772b8c8927d8d32afc8961fbc177fb723637"

type txInfo struct {
	TxHash string
	Block  string
	Age    string
	From   string
	To     string
	Value  float64
	TxFee  float64
}

func ethGet(url string) (*http.Response, []error) {
	ua := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36"
	cookies := "__cfduid=d77c0d99f010ac53bf8ac51b8398aa0591501555985; ASP.NET_SessionId=n2rakpshbewr5nymgmiaqxbx; _ga=GA1.2.430946762.1501555988; _gid=GA1.2.595151211.1502156562"
	request := gorequest.New()
	resp, _, errs := request.Get(url).Set("User-Agent", ua).Set("Cookie", cookies).End()
	fmt.Println(errs)
	return resp, errs
}

func ethDoc(doc *goquery.Document) {
	tbody := doc.Find("div.table-responsive").Find("tbody")
	trs := tbody.Find("tr")
	fmt.Println(trs.Length())
	lens := trs.Length()
	for i := 0; i < lens; i++ {
		tr := trs.Eq(i)
		tds := tr.Find("td")
		tdslen := tds.Length()
		for j := 0; j < tdslen; j++ {
			td := tds.Eq(j)
			switch j {
			case 0:
				{
					fmt.Println(td.Text())
				}
			case 1:
				{
					fmt.Println(td.Text())
				}
			case 2:
				{
					fmt.Println(td.Text())
				}
			case 3:
				{
					fmt.Println(td.Text())
				}
			case 4:
				{
					fmt.Println(td.Text())
				}
			case 5:
				{
					fmt.Println(td.Text())
				}
			case 6:
				{
					fmt.Println(td.Text())
				}

			}
		}
	}
}
func main() {
	fmt.Println("start")
	now := time.Now()
	fmt.Println(now)
	//fmt.Println(ethGet(url))
	resp, _ := ethGet(url)
	if nil != resp {
		doc, err := goquery.NewDocumentFromResponse(resp)
		if err == nil {
			ethDoc(doc)
		} else {
			fmt.Println(err)
		}
	}
	return
}
