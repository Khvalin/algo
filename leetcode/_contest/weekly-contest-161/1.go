package main

func minimumSwap(s1 string, s2 string) int {
	if len(s1) != len(s2) {
		return -1
	}

	r1, r2 := 0, 0

	x1, x2, y1, y2 := 0, 0, 0, 0

	for i := 0; i < len(s1); i++ {
		if s1[i] == 'x' {
			x1++
		} else {
			y1++
		}

		if s2[i] == 'x' {
			x2++
		} else {
			y2++
		}

		if s1[i] == 'x' && s2[i] == 'y' {
			r1++
		}

		if s1[i] == 'y' && s2[i] == 'x' {
			r2++
		}
	}

	if x1%2 != x2%2 || y1%2 != y2%2 {
		return -1
	}

	return (r1 >> 1) + (r2 >> 1) + r1%2 + r2%2
}
