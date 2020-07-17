package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
}

type HelloRequest struct {
	Name    string
	Message string
}

func (s *HelloService) SayHello(request *HelloRequest, response *string) error {
	*response += request.Name + "," + request.Message
	return nil
}

func main() {
	// 创建监听
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("failed to accept:%v", err)
	}
	rpc.RegisterName("HelloService", &HelloService{})
	// 启动服务，并将rpc服务绑定到conn上
	rpc.ServeConn(conn)
}
