package main

func main() {
	error3()
}

// 错误例子1
// fatal error: all goroutines are asleep - deadlock!
// 在main函数的协程中，向通道写入数据，而此时接收方并为准备好接收数据，导致阻塞
func error1() {
	// 创建一个无缓冲整型通道
	ch := make(chan int)
	//ch := make(chan int,0)

	// 向通道写入数据
	ch <- 1

	for c := range ch {
		println(c)
	}
}

// 错误列子2
// fatal error: all goroutines are asleep - deadlock!
// 在main函数的协程中，向通道写入数据
// for循环一直在读取数据，通道数据被读取后，下一个读取导致了阻塞
func error2() {
	// 创建一个有缓冲整型通道
	ch := make(chan int, 1)
	// 向通道写入数据
	ch <- 1

	// 加入下述代码，也会报同样的错误，因为通道缓冲已满
	//ch <- 2

	// 加入下述代码，显式关闭通道，可以修复当前函数报错
	//close(ch)

	for c := range ch {
		println(c)
	}
}

func error3() {
	ch := make(chan int)

	go pump(ch)

	println(<-ch)
}

func pump(ch chan int) {
	ch <- 1
}
