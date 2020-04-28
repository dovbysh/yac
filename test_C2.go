package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 65535)
	writer := bufio.NewWriterSize(os.Stdout, 65535)
	var len uint64
	_, err := fmt.Fscanln(reader, &len)
	if err != nil {
		panic(err)
	}
	var curr, prev int32
	prev = math.MinInt32
	for i := len; i > 0; i-- {
		_, err := fmt.Fscanln(reader, &curr)
		if err != nil {
			panic(err)
		}
		if curr < prev {
			panic(fmt.Errorf("curr < prev: %d < %d", curr, prev))
		}
		if curr > prev || i == len {
			fmt.Fprintln(writer, curr)
		}
		prev = curr
	}
	writer.Flush()
}
