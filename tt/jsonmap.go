package main

import (
	"encoding/json"
	"fmt"
)

type Msg struct {
	Dos map[string]string
}

func main() {
	msg := Msg{}
	msg.Dos = make(map[string]string)
	msg.Dos["key1"] = "value1"
	msg.Dos["key2"] = "value2"
	b, err := json.Marshal(&msg)
	fmt.Println(string(b))
	fmt.Println(err)

	f1map := make(map[string]interface{})
	f2map := make(map[string]string)
	f2map["k1"] = "v1"
	f2map["k2"] = "v2"
	f1map["f2"] = f2map
	body, err1 := json.Marshal(f1map)
	fmt.Println(string(body))
	fmt.Println(err1)

}
