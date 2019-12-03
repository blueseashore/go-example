package main

import "fmt"

var fb2 chan uint64

func generateFb2() chan uint64 {
	yield := make(chan uint64)
	var num uint64 = 0
	var num2 uint64 = 1

	go func() {
		yield <- 1
		for {
			var total = num + num2
			yield <- total
			num = num2
			num2 = total
		}
	}()
	return yield
}
func main() {
	fb2 = generateFb2()
	for i := 0; i < 10; i++ {
		fmt.Println(<-fb2)
	}
}
