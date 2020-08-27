package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	testAtomicOperation()
}

func testAtomicOperation() {
	var num uint32
	num = 100
	delta := int32(-3)
	atomic.AddUint32(&num, uint32(delta))
	fmt.Println(num)

	var num2 uint32
	num2 = 100
	delta = -3
	atomic.AddUint32(&num2, uint32(delta))
	fmt.Println(num2)

}
