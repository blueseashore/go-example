package main

// 该程序报错
// fatal error: all goroutines are asleep - deadlock! 所有的 协程（goroutines）都处于休眠（阻塞）状态
// main函数体内，创建了无缓冲通道，
// <-ch，从通道读取数据时，执行main函数的协程被阻塞，无法继续执行，从而导致了死锁
func main() {
	var ok = true
	ch := make(chan int)

	// 循环从通道里读取值，去掉循环后不会报错
	for ok {
		i := <-ch
		println("ok is %t and the counter is at %d", ok, i)
	}

	go tel(ch)
}

// 向通道写入数据
func tel(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
}
