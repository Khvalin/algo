package main

import (
	"fmt"
)

func marvelousNumber(n string) string {
	const F, S = ('4'), ('7')

	var solve func(dgts []byte) []byte
	solve = func(dgts []byte) []byte {
		L := len(dgts)
		fallback := func() []byte {
			l := L + 1
			if L%2 == 0 {
				l++
			}
			r := make([]byte, l)
			for i := 0; i < l; i++ {
				if i < l/2 {
					r[i] = F
				} else {
					r[i] = S
				}
			}

			return r
		}
		if L%2 == 1 {
			return fallback()
		}

		c4, c7 := 0, 0
		f := L + 1
		for i := 0; i < L; i++ {
			if dgts[i] <= F {
				c4++
			}
			if dgts[i] > F && dgts[i] <= S {
				c7++
			}
			if f > i && dgts[i] > S {
				f = i
			}
		}

		if f < L {
			for i := f + 1; i < L; i++ {
				dgts[i] = '0'
			}
			for f >= 0 && dgts[f] > S {
				f--
				if f >= 0 {
					dgts[f]++
				}
			}

			if f < 0 {
				dgts = append([]byte{'1'}, dgts...)
			}
			return solve(dgts)
		}

		if c7 > c4 {
			return fallback()
		}
		for i := L - 1; i >= 0; i-- {
			ch := dgts[i]
			if ch < S {
				if ch > F {
					dgts[i] = S
				} else {
					if c7 < c4 {
						dgts[i] = S
						c7++
						c4--
					} else {
						dgts[i] = F
					}
				}

			}
		}
		return dgts
	}

	return string(solve([]byte(n)))
}

func main() { /*
		fmt.Println(marvelousNumber("4587"))
		fmt.Println(marvelousNumber("7779"))
		fmt.Println(marvelousNumber("1007"))

		fmt.Println(marvelousNumber("4478"), "4747")*/
	fmt.Println(marvelousNumber("4449"), "4477")

	fmt.Println(marvelousNumber("747777"), "774447")
	fmt.Println(marvelousNumber("4849"), "7447")
	fmt.Println(marvelousNumber("1000"), "4477")
}
