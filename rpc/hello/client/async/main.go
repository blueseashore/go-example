/**
 * 异步阻塞调用示例
 */
package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type HelloRequest struct {
	Name    string
	Message string
}

func main() {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatalf("failed to dial:%v", err)
	}
	helloCall := client.Go("HelloService.SayHello", &HelloRequest{Name: "ken", Message: "is me"}, new(string), nil)

	// 其他操作

	helloCall = <-helloCall.Done

	if err := helloCall.Error; err != nil {
		log.Fatal(err)
	}
	args := helloCall.Args.(interface{})
	reply := helloCall.Reply.(*string)
	fmt.Println(args)
	// *是指针运算符 , 可以表示一个变量是指针类型 ,
	// 也可以表示一个指针变量所指向的存储单元 , 也就是这个地址所存储的值 .
	fmt.Println(*reply)
	// & 是取地址符号 , 即取得变量的地址
	fmt.Println(&reply)
}
