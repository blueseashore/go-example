package main

import (
	"context"
	pb "github.com/blueseashore/go-example/grpc/hello/hello"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received:%v", in.GetName()+","+in.GetMessage())
	return &pb.HelloResponse{Message: in.GetName() + "," + in.GetMessage()}, nil
}

func main() {
	// 创建监听
	listener, err := net.Listen("tcp", ":4567")
	// 创建监听失败
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
