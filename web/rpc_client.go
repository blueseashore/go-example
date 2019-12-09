package main

import (
	"fmt"
	"github.com/go-acme/lego/log"
	"net/rpc"
)

const serverAddress = "localhost:1234"

type Args struct {
	N, M int
}

func main() {
	client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		log.Fatal("RPC服务连接失败", err)
	}

	// 同步调用
	args := Args{7, 8}
	var reply int
	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("调用失败:", err)
	}
	fmt.Printf("%d * %d = %d \n", args.N, args.M, reply)

	// 异步调用
	var sum int
	call1 := client.Go("Args.Multiply", args, &sum, nil)
	for {
		select {
		case <-call1.Done:
			fmt.Printf("%d * %d = %d \n", args.N, args.M, sum)
			return
		default:
			fmt.Println("继续等待结果")
		}
	}
}
