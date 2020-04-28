package main

import (
	"bufio"
	"fmt"
	"os"
	"strange/pkg/d"
)

func main() {
	writer := bufio.NewWriterSize(os.Stdout, 65535)
	var len uint8
	//_, err := fmt.Scanln(&len)
	//if err != nil {
	//	panic(err)
	//}
	//if len > 11 {
	//	panic(fmt.Errorf("wrong len: %d", len))
	//}
	for len = 0; len < 5; len++ {
		d.Gen(len, writer)
		fmt.Fprintln(writer, len, "==========")
	}
	writer.Flush()
}
