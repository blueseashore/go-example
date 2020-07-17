package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
}

// 定义请求消息的结构
type HelloRequest struct {
	Name    string
	Message string
}

// 方法类型必须外部可见，所以首字母必须大写
// 方法只有2个参数
// 方法的返回值必须是error类型
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
	_ = rpc.RegisterName("HelloService", &HelloService{})
	// 启动服务，并将rpc服务绑定到conn上
	rpc.ServeConn(conn)
}
