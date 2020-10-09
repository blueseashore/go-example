package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type User struct {
	ID   int
	Name string
}

func main() {
	// uintX 转 int
	fmt.Println("uintX 转 int")
	var uNum uint = 100
	fmt.Println(int(uNum))
	var uNum8 uint8 = 100
	fmt.Println(int(uNum8))
	var uNum16 uint16 = 100
	fmt.Println(int(uNum16))
	var uNum32 uint32 = 100
	fmt.Println(int(uNum32))
	var uNum64 uint64 = 100
	fmt.Println(int(uNum64))

	fmt.Println("数字转字符串...")

	// int 转 string
	var num = 100
	fmt.Println(strconv.Itoa(num))

	// int64 转 Base进制string
	var intX int64 = 100
	fmt.Println(strconv.FormatInt(intX, 32)) // 34
	fmt.Println(strconv.FormatInt(intX, 16)) // 64
	fmt.Println(strconv.FormatInt(intX, 8))  // 144

	// string 转 int
	var s = "100"
	toInt, _ := strconv.Atoi(s)
	fmt.Println(toInt)

	// string 转 intX
	// strconv.ParseInt(数字字符,字符当前位大小,字符需转换到的位大小)，字符需转换到的位大小：0,8,16,32,64
	toInt8, _ := strconv.ParseInt(s, 10, 8)
	fmt.Println(toInt8)


	user := &User{
		ID:   1,
		Name: "hello",
	}
	// struct 转 map
	fmt.Println("struct 转 map")
	m1 := make(map[string]interface{})
	j,_ :=json.Marshal(user)
	fmt.Println(j)
	_ = json.Unmarshal(j, &m1)
	fmt.Println(m1)
}
