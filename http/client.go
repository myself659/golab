package main

import (
	"fmt"
	"strconv"
	"time"

	"crypto/sha1"
	"encoding/hex"

	"github.com/parnurzeal/gorequest"
)

var keymap = map[string]string{
	"84b19ad6235fa48eb990042476bd6084e9ceb4d8": "4a135ed7454d0e4a25228f130a93baa1a01de719",
}

func requstDo(url string) {
	r := gorequest.New()
	appkey := "84b19ad6235fa48eb990042476bd6084e9ceb4d8"

	nonce := "12345678"

	ts := strconv.Itoa(int(time.Now().Unix()))

	in := keymap[appkey] + nonce + ts
	out := sha1.Sum([]byte(in))
	hexin := out[0:len(out)]
	hexout := hex.EncodeToString(hexin)
	resp, body, errs := r.Get(url).
		Set("X-App-Key", appkey).
		Set("X-Nonce", nonce).
		Set("X-Timestamp", ts).
		Set("X-App-Checksum", hexout).End()
	fmt.Println(resp, body, errs)
}

func requstDofailed(url string) {
	r := gorequest.New()
	appkey := "84b19ad6235fa48eb990042476bd6084e9ceb4d8"

	nonce := "12345678"

	ts := strconv.Itoa(int(time.Now().Unix()))

	in := keymap[appkey] + nonce + ts
	out := sha1.Sum([]byte(in))
	hexin := out[0:len(out)]
	hexout := hex.EncodeToString(hexin)
	resp, body, errs := r.Get(url).
		Set("X-App-Key", appkey).
		Set("X-Nonce", nonce).
		Set("X-App-Checksum", hexout).End()
	fmt.Println(resp, body, errs)
}

func main() {
	url := "http://localhost:8080/v1/normal"
	requstDo(url)
	requstDofailed(url)
}
