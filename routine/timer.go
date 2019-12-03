package main

import (
	"fmt"
	"time"
)

// 定时器
func main() {
	// 创建一个定时器，每隔 1e8发送一次
	tick := time.Tick(1e8)
	// 创建一个定时器，隔 5e8 发送一次，只发送一次
	boom := time.After(5e8)

	for {
		select {
		case i:= <-tick:
			fmt.Println(i)
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom.")
			return
		default:
			fmt.Println("default.")
			time.Sleep(5e7)
		}
	}
}
