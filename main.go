package main

import (
	"fmt"
	"reflect"
)

func test(kv [2]string) {
	fmt.Println(reflect.TypeOf(kv).Name())
}

func main() {
	var kv [2]interface{}
	kv[0] ="a"
	kv[1] ="a"
	fmt.Println(kv)
	kv2 :=make(map[int]string,2)
	kv2[0] = "a"
	kv2[1] = "a"
	kv2[2] = "a"
	kv2[3] = "a"
	kv2[4] = "a"
	fmt.Println(kv2)
}
