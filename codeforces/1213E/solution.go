package main

import (
	"fmt"
)

const empty = '!'

func solve(n uint32, s1, s2 []rune) string {
	const ABC = "abc"

	l := int64(3 * n)

	count := map[rune]uint32{
		'a': n,
		'b': n,
		'c': n,
	}

	r := make([]rune, l)
	bl := map[rune]map[rune]bool{}

	for _, ch := range [...]rune{'a', 'b', 'c', empty} {
		bl[ch] = map[rune]bool{}
	}

	bl[s1[0]][s1[1]] = true
	bl[s2[0]][s2[1]] = true

	getOptions := func(min, prev rune) []rune {
		//TODO: add order
		r := make([]rune, 0, 3)

		for _, ch := range ABC {
			if ch >= min && count[ch] > 0 && !bl[prev][ch] {
				r = append(r, ch)
			}
		}

		return r
	}

	for i := int64(0); i >= 0 && i < l; {
		prev, s := empty, empty

		if i > 0 {
			prev = r[i-1]
		}

		if r[i] >= 'a' {
			s = r[i] + 1
		}

		found := false
		options := getOptions(s, prev)

		t := r[i]
		for _, ch := range options {
			t = ch
			found = true
			break
		}

		if found {
			count[t]--
			r[i] = t
			i++
		} else {
			r[i] = empty
			count[t]++
			i--
		}

		fmt.Println(string(r), count)
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
