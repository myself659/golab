package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"sync"

	"github.com/myself659/dxproxy"
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

var wg sync.WaitGroup

func genid() {
	// 11位
	for i := 55000836096; i < 55000000000; i++ {
		s := strconv.Itoa(i)
		chids <- s
		wg.Add(1)
	}
	close(chids)
}

func iddo(id string) {
	count := 300
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
	wg.Done()
}

func idsdo() {
	count := 300

	for {
		id, ok := <-chids

		fmt.Println("id:", id, ok)
		url := genurl(id, count)
		fmt.Println(url)
		//resp, err := http.Get(url)
		resp, err := dx.Get(url)
		if err != nil {
			fmt.Println(url, err)
			return
		}

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
		resp.Body.Close()
		wg.Done()
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
	name := "id-712-56.txt"
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

var dx *dxproxy.ProxyPool

func main() {
	dx = dxproxy.NewProxyPool()

	<-time.After(10 * time.Second)
	//go dispatchid()
	go validids()
	go genid()
	for i := 0; i < 600; i++ {
		go idsdo()
		<-time.After(10 * time.Millisecond)
	}
	wg.Wait()

	<-time.After(60 * time.Second)

}
