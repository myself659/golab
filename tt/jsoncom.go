package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := `{"Name":"20161122-bcconnect-161124124308.pdf","Url":"http://116.62.188.219:8090/upload-file/test30/2-20161122-bcconnect-161124124308.pdf","FileHash":"f92a196f76c342e78f8a8b85a95d7383"}`
	f := fileBC{}
	jsonDecode(s, &f)
	fmt.Println(f)
	c := comFile{}
	jsonDecode(s, &c.Bid)
	fmt.Println(c.Bid)
}

type fileBC struct {
	FileName string `json:"Name"`
	FilePath string `json:"Url"`
	FileHash string
}

type comFile struct {
	Bid fileBC
}

func jsonDecode(s string, v interface{}) {
	b := []byte(s)
	err := json.Unmarshal(b, v)
	if err != nil {
		fmt.Println("jsonDecode:", err, s)
	}
}
