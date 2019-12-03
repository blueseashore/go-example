package main

import "fmt"

var resume chan int

// 惰性生成器
func integers() chan int {
	yield := make(chan int)
	count := 1
	go func() {
		for {
			yield <- count * count
			count++
		}
	}()
	return yield
}

func generateInteger() int {
	return <-resume
}

func main() {
	resume = integers()
	for i := 0; i < 10; i++ {
		fmt.Println(generateInteger())
	}
	fmt.Println(<-resume)
}
