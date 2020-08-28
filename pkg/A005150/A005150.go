package A005150

import "fmt"

// https://oeis.org/A005150
func Say(a string) string {
	r := ""
	i := 0
	aa := []rune(a)
	c := 1
	for i < len(aa) {
		d := aa[i]
		if i+1 < len(aa) && d == aa[i+1] {
			c++
			i++
			continue
		}
		r = fmt.Sprintf("%s%d%s", r, c, string(d))
		c = 1
		i++
	}
	return r
}
