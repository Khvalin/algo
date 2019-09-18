package main

import (
	"fmt"
)

func solve(n uint32, s1, s2 []rune) string {
	const ABC = "abc"

	l := int64(3 * n)

	count := map[rune]uint32{
		'a': n,
		'b': n,
		'c': n,
	}

	getOptions := func(min rune) []rune {
		//TODO: add order
		r := make([]rune, 0, 3)

		for _, ch := range ABC {
			if ch >= min && count[ch] > 0 {
				r = append(r, ch)
			}
		}

		return r
	}

	r := make([]rune, l)
	bl := map[rune]map[rune]bool{}

	for _, ch := range [...]rune{'a', 'b', 'c', ' '} {
		bl[ch] = map[rune]bool{}
	}

	bl[s1[0]][s1[1]] = true
	bl[s2[0]][s2[1]] = true

	prev := ' '
	for i := int64(0); i > 0 && i < l; {
		s := 'a'
		if i > 0 {
			prev = r[i-1]
			if r[i] >= 'a' {
				s = r[i] + 1
			}
		}

		found := false
		options := getOptions(s)

		t := r[i]
		for _, ch := range options {
			if !bl[prev][ch] {
				t = ch
				found = true
			}
		}

		if found {
			i++
			count[t]++
			r[i] = t
		} else {
			count[t]--
			i--
		}
	}

	return string(r)
}

func main() {
	var n uint32
	var s1, s2 string

	fmt.Scan(&n)
	fmt.Scan(&s1, &s2)

	res := solve(n, ([]rune)(s1), ([]rune)(s2))

	fmt.Println(res)
}
