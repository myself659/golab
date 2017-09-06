package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type fileDesc struct {
	Login_status bool       `json:"login_status"` // 是否登陆
	Message      string     `json:"message"`      // 消息
	Data         []dataDesc `json:"data"`         // 具体数据，获取的重点
}

type dataDesc struct {
	Title           string `json:"title"`           // 标题
	Go_detail_count int    `json:"go_detail_count"` // 阅读数
	Comments_count  int    `json:"comments_count"`  // 评论数
	Display_url     string `json:"display_url"`     // url
	Behot_time      int64  `json:"behot_time"`      // 发布时间
	Chinese_tag     string `json:"chinese_tag"`     // 中文标签
	Source          string `json:"source"`          // 来源，作者
}

func prasefile(name string) {
	fd, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fd.Close()
	b, err := ioutil.ReadAll(fd)
	if err != nil {
		fmt.Println(err)
		return
	}
	fdesc := fileDesc{}
	err = json.Unmarshal(b, &fdesc)
	fmt.Println(err)
	fmt.Println(fdesc)
}

func main() {
	name := "/Users/eric/go/src/golab/ttiao/data/3435296611-300-2017-07-05.txt"
	prasefile(name)
}
