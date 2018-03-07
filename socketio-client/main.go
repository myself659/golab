package main

import (
	"fmt"
	"time"

	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

func main() {
	host := "116.62.148.208"
	port := 3380

	clientDo(host, port)
	<-time.After(60 * time.Second)
}

func clientDo(host string, port int) (*gosocketio.Client, error) {
	c, err := gosocketio.Dial(
		gosocketio.GetUrl(host, port, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		return nil, err
	}

	err = c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
		go func() {
			fmt.Errorf("Disconnected")
		}()
	})
	if err != nil {
		return nil, err
	}
	err = c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		fmt.Println("connect ok")
	})
	if err != nil {
		return nil, err
	}
	err = c.On("msg", func(h *gosocketio.Channel, msg interface{}) {
		fmt.Println("msg:", msg)
	})
	if err != nil {
		return nil, err
	}

	return c, nil

}
