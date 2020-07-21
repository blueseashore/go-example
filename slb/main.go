package main

import (
	"fmt"
	"math/rand"
)

func init() {
	rand.Seed(10)
}

func shuffle1(slice []int) {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
}

func shuffle2(indexes []int) {
	for i := len(indexes); i > 0; i-- {
		lastIdx := i - 1
		idx := rand.Intn(i)
		indexes[lastIdx], indexes[idx] = indexes[idx], indexes[lastIdx]
	}
}

func main() {
	cnt1 := make(map[int]int);
	for i := 0; i < 100000; i++ {
		var sl = []int{0, 1, 2, 3, 4, 5, 6}
		shuffle1(sl)
		cnt1[sl[0]]++

	}
	fmt.Println(cnt1)
	cnt2 := make(map[int]int)
	for i := 0; i < 100000; i++ {
		var sl = []int{0, 1, 2, 3, 4, 5, 6}
		shuffle2(sl)
		cnt2[sl[0]]++
	}

	fmt.Println(cnt2)
}
