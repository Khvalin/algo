package main

import (
	"fmt"
)

func marvelousNumber(n string) string {
	const F, S, Z = '4', '7', '0'

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

		inc := func(f int, clause func() bool) bool {
			for i := f + 1; i < L; i++ {
				dgts[i] = Z
			}

			for f >= 0 && ((clause != nil) && clause() || dgts[f] > S) {
				if dgts[f] > F {
					c4++
					c7--
				}
				dgts[f] = Z
				f--
				if f >= 0 {
					if dgts[f] == F {
						c4--
						c7++
					}
					dgts[f]++
				}
			}

			of := f < 0
			if of {
				dgts = append([]byte{'1'}, dgts...)
			}
			return of
		}
		///

		pr, f := false, L+1
		for i, ch := range dgts {
			if !pr && ch != F && ch < S {
				pr = !inc(i, nil)
			}

			if ch <= F {
				c4++
			}
			if ch > F && ch <= S {
				c7++
			}
			if f > i && ch > S {
				f = i
			}
		}

		if f < L {
			inc(f, nil)

			return solve(dgts)
		}

		if c7 > c4 {
			clause := func() bool { return c7 > c4 }

			if inc(L-1, clause) {
				return fallback()
			}
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

		if c4 != c7 {
			return fallback()
		}
		return dgts
	}

	return string(solve([]byte(n)))
}

func main() {

	fmt.Println(marvelousNumber("4587"))
	fmt.Println(marvelousNumber("7779"))
	fmt.Println(marvelousNumber("1007"))

	fmt.Println(marvelousNumber("4478"), "4747")
	fmt.Println("4449 => ", marvelousNumber("4449"), "4477")

	fmt.Println(marvelousNumber("747777"), "774447")
	fmt.Println(marvelousNumber("4849"), "7447")
	fmt.Println(marvelousNumber("1000"), "4477")

	fmt.Println(marvelousNumber("444777"), "444777")
	fmt.Println(marvelousNumber("474747777777444449"), "474774444444777777")

	fmt.Println(marvelousNumber("7777744445"), "444444777777")

	fmt.Println(marvelousNumber("70070077"), "74444777")
}
