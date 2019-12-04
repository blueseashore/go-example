// 链式操作

package main

import (
	"flag"
	"fmt"
)

// 返回一个int指针
var ngoroutine = flag.Int("n", 100000, "How many goroutines")

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	flag.Parse()
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost

	// *ngoroutine 对指针取值
	for i := 0; i < *ngoroutine; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}

	right <- 0
	fmt.Println(<-leftmost)
}
