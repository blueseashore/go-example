package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type C struct {
	sync.RWMutex
	items map[string]string
}

func R(c *C) {
	c.Lock()
	fmt.Println(c.items["name"], c.items["age"], c.items["sex"])
	c.Unlock()
}

func W(c *C, i int) {
	c.Lock()
	c.items["name"] = "zz"
	c.items["age"] = strconv.Itoa(rand.Intn(100))
	c.items["sex"] = "female"
	c.Unlock()
}
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("单一")
		}
	}()
	mp := make(map[string]string)
	mp["name"] = "ken"
	mp["age"] = "20"
	mp["sex"] = "man"
	c := &C{items: mp}

	for i := 1; i < 10000; i++ {
		go W(c, i)
		go R(c)

	}
	time.Sleep(time.Second)
}
