package main

func scariestMask(mask []string) (res int) {
	diff := func(s string) (r int) {
		i, j := 0, len(s)-1
		for i < j {
			if s[i] != s[j] {
				r++
			}

			i++
			j--
		}

		return
	}

	for _, s := range mask {
		res += diff(s)
	}
	return
}
