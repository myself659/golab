package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func genKeyPair() (string, string) {
	u1 := uuid.New()

	in := u1[0:len(u1)]
	tb := []byte(string(time.Now().Unix()))
	in = append(in, tb...)
	// appkey =  uuid +  timestamp
	out := sha1.Sum(in)
	hexin := out[0:len(out)]
	appkey := hex.EncodeToString(hexin)

	sha1in := append(hexin, []byte(string(time.Now().Unix()))...)
	sha1in = append(sha1in, []byte(stringWithCharset(16, charset))...)
	// secretkey = appkey + timestamp + rand str
	sha1out := sha1.Sum(sha1in)
	hexsecretin := sha1out[0:len(sha1out)]
	hexsecretout := hex.EncodeToString(hexsecretin)

	return appkey, hexsecretout
}

func main() {
	fmt.Println(genKeyPair())
}
