package funcs

func countStrange(s1, s2 string) int64 {
	counter := int64(0)
	ls2 := len(s2)
	ls1 := len(s1)
	var b1, b2 byte
	for i := 0; i < ls2; i++ {
		b2 = s2[i]
		for j := 0; j < ls1; j++ {
			b1 = s1[j]
			if b2 == b1 {
				counter++
			}
		}
	}
	return counter
}
