package main

import (
	"fmt"
	"math"
)

func main() {
	var len uint64
	_, err := fmt.Scanln(&len)
	if err != nil {
		panic(err)
	}
	var curr, prev int32
	prev = math.MinInt32
	for i := len; i > 0; i-- {
		_, err := fmt.Scanln(&curr)
		if err != nil {
			panic(err)
		}
		if curr < prev {
			panic(fmt.Errorf("curr < prev: %d < %d", curr, prev))
		}
		if curr > prev || i == len {
			fmt.Println(curr)
		}
		prev = curr
	}
}
