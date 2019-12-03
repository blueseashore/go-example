package main

import "fmt"

// 创建一个数组接收斐波那契数列的值
var fb [] uint64

func generateFb() uint64 {
	var num uint64 = 1
	if len(fb) > 1 {
		num = fb[len(fb)-2] + fb[len(fb)-1]
	}
	fb = append(fb, num)
	return num
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(generateFb())
	}
}
