package main

import (
	"fmt"
	"time"
)

const (
	// 10 * 2^20 = 10 *1024 * 1024,10M
	AvailableMemory = 10 << 20

	// 10KB
	AverageMemoryPerRequest = 10 << 10

	MAXREQUEST = AvailableMemory / AverageMemoryPerRequest
)

var sem = make(chan int, MAXREQUEST)

type Request struct {
	a, b   int
	reply chan int
}

func process(r *Request) {
	fmt.Println(time.Now())
}

func handle(r *Request) {
	process(r)
	<-sem
}

func Server(queue chan *Request) {
	for {
		sem <- 1
		request := <-queue
		go handle(request)
	}
}

func main() {
	queue := make(chan *Request)
	fmt.Println(1)
	go Server(queue)
	time.Sleep(1e9)
}
