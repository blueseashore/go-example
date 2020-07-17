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
	var response string
	err = client.Call("HelloService.SayHello", &HelloRequest{Name: "ken", Message: "is me"}, &response)
	if err != nil {
		log.Fatalf("failed to call:%v", err)
	}
	fmt.Println(response)
}
