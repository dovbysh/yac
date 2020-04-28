package main

import (
	"bufio"
	"fmt"
	"os"
)

func resP(s1, s2 *map[byte]uint32) {
	if len(*s1) != len(*s2) {
		fmt.Println("0")
		return
	} else {
		for k, v := range *s1 {
			zz, ok := (*s2)[k]
			if !ok {
				fmt.Println("0")
				return
			}
			if zz != v {
				fmt.Println("0")
				return
			}
		}
		fmt.Println("1")
	}
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 65535)
	s1 := make(map[byte]uint32)
	s2 := make(map[byte]uint32)
	var s *map[byte]uint32
	var b byte
	var e error
	s = &s1
	var c, cc uint32
	for {
		b, e = reader.ReadByte()
		if e != nil {
			fmt.Println(e)
			break
		}
		c++
		if c > 100000 {
			panic(fmt.Errorf("too long"))
		}
		if b == '\n' {
			c = 0
			cc++
			s = &s2
			if cc > 1 {
				break
			}
			continue
		}
		if b < 'a' || b > 'z' {
			panic(fmt.Errorf("invalid char"))
		}
		(*s)[b]++
	}
	resP(&s1, &s2)
}
