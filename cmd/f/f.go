package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 65535)
	//writer := bufio.NewWriterSize(os.Stdout, 65535)
	var slen, alen uint32
	_, err := fmt.Fscanln(reader, &slen)
	if err != nil {
		panic(err)
	}
	m := make([][]byte, slen, slen)
	res := make([]byte, 0)
	var i, j uint32
	var b byte
	var curMin uint32
	for i = 0; i < slen; i++ {
		s, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		st := strings.Split(s, " ")
		if len(st) == 0 {
			panic("zerro")
			//m[i] = make([]byte, 0, 0)
			//continue
		}

		_, err = fmt.Sscan(st[0], &alen)
		if err != nil {
			panic(err)
		}
		if int(alen) != len(st)-1 {
			panic("invalid format")
		}
		m[i] = make([]byte, alen, alen)
		for j = 1; j <= alen; j++ {
			_, err := fmt.Sscan(st[j], &b)
			if err != nil {
				panic(err)
			}
			m[i][j-1] = b
		}
	}
	fmt.Println(m)
	for {
		slen = uint32(len(m))
		if slen == 0 {
			break
		}
		for i = 0; i < slen; i++ {
			if m[i][0] < m[curMin][0] {
				curMin = i
			}
		}
		res = append(res, m[curMin][0])
		if len(m[curMin]) > 0 {
			m[curMin] = m[curMin][1:]
		}
		if len(m[curMin]) == 0 {
			copy(m[curMin:], m[int(curMin)+1:])
			m[len(m)-1] = nil
			m = m[:len(m)-1]
		}
		fmt.Println(curMin)
		curMin = 0
	}
	fmt.Println(res)
}
