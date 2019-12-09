package main

import (
	"fmt"
	"github.com/go-acme/lego/log"
	"golang.org/x/net/websocket"
	"net/http"
)

func server(ws *websocket.Conn) {
	fmt.Printf("新连接")
	fmt.Println(ws)
	buf := make([]byte, 100)
	for {
		if _, err := ws.Read(buf); err != nil {
			fmt.Printf("%s", err.Error())
			break
		}
		fmt.Println(string(buf))
		ws.Write([]byte("i know"))
	}
	fmt.Printf("=>关闭连接")
	ws.Close()
}

func main() {
	http.Handle("/websocket", websocket.Handler(server))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("启动服务失败", err.Error())
	}
}
