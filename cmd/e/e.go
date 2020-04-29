package main

import (
	"fmt"
	"io"
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
	s1 := make(map[byte]uint32)
	s2 := make(map[byte]uint32)
	var s *map[byte]uint32
	var b byte
	var e error
	var n int
	s = &s1
	var c, cc uint32

	buf := make([]byte, 4096)
	for {
		if cc > 1 {
			break
		}
		n, e = os.Stdin.Read(buf)
		if e == io.EOF {
			break
		}
		for i := 0; i < n; i++ {
			b = buf[i]
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
	}
	resP(&s1, &s2)
}
