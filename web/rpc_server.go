package main

import (
	"github.com/go-acme/lego/log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type Args struct {
	N, M int
}

func (t *Args) Multiply(args *Args, reply *int) error {
	*reply = args.N * args.M
	return nil
}

func main() {
	// 创建对象
	calc := new(Args)
	// 注册服务
	rpc.Register(calc)
	rpc.HandleHTTP()

	// 创建监听
	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("启动RPC服务失败", err)
	}
	go http.Serve(listener, nil)

	os.Stdin.Read(make([]byte, 1))
	// sleep 1000秒
	//time.Sleep(1000 * time.Second)
}
