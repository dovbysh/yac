package main

import "fmt"

func main() {
	var len uint64
	_, err := fmt.Scanln(&len)
	if err != nil {
		panic(err)
	}
	var curr uint8
	var count, maxCount uint64
	for i := len; i > 0; i-- {
		_, err := fmt.Scanln(&curr)
		if err != nil {
			panic(err)
		}
		if curr > 1 {
			panic(fmt.Errorf("out of range: %d", curr))
		}
		if curr == 0 {
			if count > maxCount {
				maxCount = count
			}
			count = 0
			continue
		}
		count++
	}
	if count > maxCount {
		maxCount = count
	}
	fmt.Println(maxCount)
}
