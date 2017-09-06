package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type bindPara struct {
	Name string `json:"Uuid"`
	Cert string `json:"Cert"`
}

func postjsonDo(url string, i interface{}, rsp interface{}) error {
	b, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(b))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	/*
		reqdup := req
		rb, err := ioutil.ReadAll(reqdup.Body)
		fmt.Println("do:", string(rb))
	*/
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("postjsonDo response Status:", resp.Status)
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, rsp)
	if err != nil {
		return err
	}
	fmt.Println("postjsonDo:", url, rsp, i)
	return nil
}

type baseBCRsp struct {
	Status string `json:"status"`
	TxID   string `json:"tx_id"`
}

func userBCBind(name string) baseBCRsp {
	var rsp baseBCRsp
	bind := bindPara{Name: name, Cert: "helo"}
	url := "http://192.168.20.190:3000/api/chain/user/bind"
	err := postjsonDo(url, bind, &rsp)
	if err != nil {
		fmt.Println("userBCBind:", err)
	}

	return rsp
}

func postjson(url string, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	fmt.Println(req.Body)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}

func main() {
	userBCBind("g12345")
	/*
		url := "http://192.168.20.190:3000/api/chain/user/bind"

		bind := bindPara{Name: "testddd", Cert: "ddd"}
		postjson(url, bind)
	*/
}
