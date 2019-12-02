package main

func main() {
	var i int
	var ok = true
	ch := make(chan int)

	go tel2(ch)
	for ok {
		// channel关闭后，ok=false
		if i, ok = <-ch; ok {
			println("ok is %t and the counter is at %d", ok, i)
		} else {
			println("finished")
		}
	}
}

// 向通道写入数据
func tel2(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	// 显式关闭通道
	// 注释下方代码后，报错：fatal error: all goroutines are asleep - deadlock!
	close(ch)
}
