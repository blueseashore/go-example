package main

import (
	"fmt"
	"github.com/go-acme/lego/log"
	"golang.org/x/net/websocket"
	"time"
)

func main() {
	ws, err := websocket.Dial("ws://localhost:12345/websocket", "", "http://localhost/")
	if err != nil {
		log.Fatal("握手失败", err.Error())
	}
	ws.Write([]byte("hello is me"))
	go readFromServer(ws)
	time.Sleep(5 * time.Second)
	ws.Close()
}

func readFromServer(ws *websocket.Conn) {
	buf := make([]byte, 1000)
	for {
		if _, err := ws.Read(buf); err != nil {
			fmt.Printf("%s\n", err.Error())
			break
		}
		fmt.Println(string(buf))
	}
}
