package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:50001")
	if err != nil {
		fmt.Println("创建连接失败：", err.Error())
		return // 终止程序
	}
	// 创建命令行读取对象
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入你的名字")
	clientName, _ := inputReader.ReadString('\n')
	name := strings.Trim(clientName, "\r\n")

	for {
		fmt.Println("请输入需要发送的心情，输入大写Q退出")
		input, _ := inputReader.ReadString('\n')
		message := strings.Trim(input, "\r\n")

		if message == "Q" {
			fmt.Println("退出")
			return
		}
		_, err = conn.Write([]byte(name + " 说：" + message))
		if err != nil {
			fmt.Println("发送失败:", err.Error())
		}else{
			fmt.Println("success")
		}
	}
}
