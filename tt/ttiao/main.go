package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func genurl(id string, count int) string {
	sc := strconv.Itoa(count)
	url := "http://www.toutiao.com/c/user/article/?page_type=1&user_id=" + id + "&count=" + sc
	return url
}

func save(id string, count int, b []byte) {
	fname := id + "-" + strconv.Itoa(count) + "-" + strings.Split(time.Now().String(), " ")[0] + ".txt"
	fd, err := os.Create("data/" + fname)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = fd.Write(b)
	fmt.Println(err)

	return
}

func dispatchid() {
	// line by line read
	defer close(chids)
	fr, err := os.Open("ids.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fr.Close()
	scanner := bufio.NewScanner(fr)
	for scanner.Scan() {
		s := scanner.Text()
		s = strings.TrimSpace(s)
		if s != "" {
			//fmt.Println(s)
			chids <- s
		}
	}
	fmt.Println("end  disptatch id")

}

func idsdo() {

	count := 300
	for {
		id, ok := <-chids
		if ok == false {
			chexit <- true
			return
		}
		fmt.Println("id:", id)
		url := genurl(id, count)
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(url, err)
			return
		}
		defer resp.Body.Close()

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(url, err)

		} else {
			//fmt.Println(string(b))
			if len(b) > 200 {
				save(id, count, b)
			}
		}

	}

}

var chids = make(chan string, 16)
var chexit = make(chan bool)

func main() {
	//fmt.Println("start")
	/*
		var url string
		url = "http://www.toutiao.com/c/user/article/?page_type=1&user_id=53324741611"
		url = "http://www.toutiao.com/c/user/article/?page_type=1&user_id=53324741611&max_behot_time=1497275211&count=20"
		url = "http://www.toutiao.com/c/user/article/?page_type=1&user_id=53324741611&count=20"
		url = "http://www.toutiao.com/c/user/article/?page_type=1&user_id=53324741611&count=400"
		// http://www.toutiao.com/c/user/51558403490/#mid=51558403490
		url = "http://www.toutiao.com/c/user/article/?page_type=1&user_id=51558403490&count=2000" // failed
		url = "http://www.toutiao.com/c/user/article/?page_type=1&user_id=51558403490&count=400"

		id = "51558403490"
		count = 300
		id = "51559696103"
		id = "54570120686"
		id = "53324741611"
		id = "59673343446"
		id = "60196232057"
		id = "5905739249"
		id = "53117304178"
	*/
	go dispatchid()
	go idsdo()
	<-chexit
	<-time.After(10 * time.Second)

}
