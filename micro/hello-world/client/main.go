package main

import (
	"context"
	"fmt"
	"github.com/blueseashore/go-example/micro/hello-world/proto"
	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()

	greeter := proto.NewGreeterService("greeter", service.Client())

	rsp, err := greeter.Hello(context.TODO(), &proto.Request{Name: "Ken"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp.Msg)
}
