package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"encoding/json"

	proxy "github.com/myself659/xiciproxy"
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
	//持久化
	// 11位
	var start = 55900000000
	name := "idoffset.txt"
	fd, err := os.OpenFile(name, os.O_SYNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		fd, err = os.Create(name)
		if err != nil {
			fmt.Println("genid:", err)
		}
	} else {
		b, err := ioutil.ReadAll(fd)
		fmt.Println("genids:", string(b), err)
		start, err = strconv.Atoi(string(b))
		fmt.Println("genids:", start, err)
	}
	if start == 0 {
		start = 55900000000
	}

	for i := start; i > 51000000000; i-- {
		s := strconv.Itoa(i)
		chids <- s
		fd.Truncate(0)
		fd.Seek(0, 0)
		fd.Write([]byte(s))
		fd.Sync()
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
		fmt.Println("idsdo:", id)
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
			length := len(b)
			fdesc := fileDesc{}
			err := json.Unmarshal(b, &fdesc)
			fmt.Println("doid:", id, err, length)
			if err == nil && length > 200 {
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

type fileDesc struct {
	Login_status bool       `json:"login_status"`
	Message      string     `json:"message"`
	Data         []dataDesc `json:"data"`
}

type dataDesc struct {
	Title           string `json:"title"`
	Go_detail_count int    `json:"go_detail_count"`
	Comments_count  int    `json:"comments_count"`
	Display_url     string `json:"display_url"`
	Behot_time      int64  `json:"behot_time"`
	Chinese_tag     string `json:"chinese_tag"`
	Source          string `json:"source"`
}

func validids() {
	name := "id-11-6.txt"
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

var dx *proxy.ProxyPool

func main() {
	dx = proxy.NewProxyPool()

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
