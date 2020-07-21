package main

import (
	"context"
	"fmt"
	"github.com/blueseashore/go-example/micro/pb"
	"github.com/micro/go-micro/v2"
)

type Greeter struct {
}

func (g *Greeter) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Greeting = "Hello" + req.Name
	return nil
}

func main() {
	service := micro.NewService(micro.Name("greeter"))

	service.Init()

	_ = pb.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
