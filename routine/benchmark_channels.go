package main

import (
	"fmt"
	"testing"
)

func main()  {
	fmt.Println("sync",testing.Benchmark(benchmarkChannelSync).String())
}

func benchmarkChannelSync(b *testing.B) {
	ch := make(chan int)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	for _ = range ch {

	}
}

