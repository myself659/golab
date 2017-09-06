package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 不支持小写
type emd struct {
	To string `json:"to,omitempty"`
}

type resp struct {
	Code   float64  `json:"code,omitempty"`
	Msg    string   `json:"msg,omitempty"`
	Ignore int      `json:"-,omitempty"`
	Data   emd      `json:"data,omitempty"`
	Files  []string `json:"files,omitempty"`
}

type test struct {
	Code int
	Data string
}

/* 高效的处理过程 */

func main() {
	a := resp{Code: 1, Msg: "hello", Ignore: 35,
		Data: emd{
			To: "world",
		},
		Files: []string{"a", "b"},
	}
	fmt.Println(a)
	b, err := json.Marshal(a)
	if err == nil {
		os.Stdout.Write(b)
	}
	s := `{"Code":1,"Msg":"hello","Data":{"To":"world"},"Files":["a","b"]}`
	bs := []byte(s)
	ra := resp{}
	err = json.Unmarshal(bs, &ra)
	fmt.Println(err, ra)
	ss := `{"Code":0,"Data":"{\"hash\":\"xx\",\"time\":\"2017\",\"str\":\"url://\"}"}`
	t := test{}
	err = json.Unmarshal([]byte(ss), &t)
	fmt.Println(err, t)
}
