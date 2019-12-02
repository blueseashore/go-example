package main

// 使用select切换携程的例子
import "time"

func main() {
	// 创建2个通道
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 开启协程，向通道写入数据
	go pump1(ch1)
	go pump2(ch2)
	// 开启协程，从通道取出数据
	go suck(ch1, ch2)

	// 程序睡眠1秒，等到协程启动成功
	time.Sleep(1e9)
}

// 向通道写入数据
func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

// 向通道写入数据
func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

// 从通道取数据
func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			println("Received on channel 1:%d", v)
		case v := <-ch2:
			println("Received on channel 2:%d", v)
		}
	}
}
