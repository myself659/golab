package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/myself659/xiciproxy"
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

func genid() {
	// 11位
	for i := 50000000001; i < 60000000000; i++ {
		s := strconv.Itoa(i)
		chids <- s
	}
	close(chids)
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
		//resp, err := http.Get(url)
		resp, err := dx.Get(url)
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
				// save id
				chvalidids <- id
			}
		}

	}

}

var chids = make(chan string, 16)
var chvalidids = make(chan string, 16)
var chexit = make(chan bool)

func validids() {
	name := "id-11-5.txt"
	fd, err := os.Create(name)
	if err != nil {
		fmt.Println("validids:", err)
		return
	}
	defer fd.Close()
	for {
		sid := <-chvalidids
		sid = sid + "\r\n"
		fd.Write([]byte(sid))

	}
}

var dx *xiciproxy.ProxyPool

func main() {
	dx = xiciproxy.NewProxyPool()

	<-time.After(10 * time.Second)
	//go dispatchid()
	go validids()
	go genid()
	for i := 0; i < 3; i++ {
		go idsdo()
	}

	<-chexit
	<-time.After(60 * time.Second)

}
