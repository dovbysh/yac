package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func wrapB(s []byte) {
	s[0] = '('
	s[len(s)-1] = ')'
}
func appendB(s []byte) {
	s[len(s)-2] = '('
	s[len(s)-1] = ')'
}
func prepB(s []byte) {
	s[0] = '('
	s[1] = ')'
}

type f func(b []byte)

var br [][]f

func init() {
	br = make([][]f, 12, 12)
	br[0] = []f{func([]byte) {}}
	br[1] = []f{func(b []byte) {
		wrapB(b)
	}}
	br[2] = []f{
		func(b []byte) {
			wrapB(b[1:3])
			wrapB(b)
		},
		func(b []byte) {
			br[1][0](b[0:2])
			appendB(b)
		},
	}
	for i := 3; i < 5; i++ {
		l := i
		br[l] = make([]f, 0)
		for _, ff := range br[l-1] {
			fff := ff
			br[l] = append(br[l],
				func(b []byte) {
					fff(b[1 : l*2-1])
					wrapB(b)
				},
			)
		}
		br[l] = append(br[l], func(b []byte) {
			br[l-1][0](b[0 : l*2-2])
			appendB(b)
		},
		)
		for _, ff := range br[l-1] {
			fff := ff
			br[l] = append(br[l], func(b []byte) {
				fff(b[2:])
				prepB(b)
			},
			)
		}
	}
}

//   ((()))
//   (()())
//   (())()
//   ()(())
//   ()()()
//

func gen(len uint8, w io.Writer) {
	b := make([]byte, len*2, len*2)
	for _, ff := range br[len] {
		ff(b)
		writeB(w, b)
	}
}

func writeB(w io.Writer, b []byte) {
	w.Write(b)
	w.Write([]byte("\n"))
}
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
		gen(len, writer)
		fmt.Fprintln(writer, "==========")
	}
	writer.Flush()
}
