package main

import (
	"context"
	"fmt"
	pb "github.com/blueseashore/go-example/grpc/hello/hello"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

const (
	defaultName    = "jack"
	defaultMessage = "hello"
)

func main() {
	conn, err := grpc.Dial("localhost:4567", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to dial,%v", err)
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	name := defaultName
	message := defaultMessage
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	if len(os.Args) > 2 {
		message = os.Args[2]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name, Message: message})
	if err != nil {
		log.Fatalf("failed to SayHello: %v", err)
	}
	fmt.Println(r.GetMessage())
}
