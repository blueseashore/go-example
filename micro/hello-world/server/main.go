package main

import (
	"context"
	"fmt"
	"github.com/blueseashore/go-example/micro/hello-world/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry/etcd"
	_ "github.com/micro/go-plugins/transport/rabbitmq"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Msg = "Name:" + req.Name
	return nil
}

func main() {
	registry := etcd.NewRegistry()

	service := micro.NewService(
		micro.Name("greeter"),
		micro.Registry(registry),
	)

	service.Init()

	_ = proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err.Error())
	}
}
