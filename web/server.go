package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("启动服务...")
	// 创建监听
	listener, err := net.Listen("tcp", "localhost:50001")
	if err != nil {
		fmt.Println("监听错误:", err.Error())
		return // 终止程序
	}

	// 监听并接受来自客户端的连接
	for {
		fmt.Println("接收请求")
		conn, err := listener.Accept()
		fmt.Println(err)
		if err != nil {
			fmt.Println("接收错误:", err.Error())
			return // 终止程序
		}
		fmt.Println("处理请求")
		// 处理请求连接
		go doServerStuff(conn)
	}
}

// 处理请求连接
func doServerStuff(conn net.Conn) {
	fmt.Println("进行读取")
	for {
		buf := make([]byte, 512)
		lens, err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取错误:", err.Error())
			return // 终止程序
		}
		fmt.Println("读取的数据是：", string(buf[:lens]))
	}
}
