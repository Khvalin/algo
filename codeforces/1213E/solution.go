package main

import (
	"fmt"
)

const empty = '!'
const nmax = 4

func sample(n uint32, bl map[rune]map[rune]bool, abc string) (bool, string, uint32) {
	l := int64(3 * n)

	r := make([]rune, l)

	count := map[rune]uint32{
		'a': n,
		'b': n,
		'c': n,
	}

	order := map[rune]int{}

	order[empty] = 0
	for i, ch := range abc {
		order[ch] = i + 1
	}

	getOptions := func(min, prev rune) rune {
		//TODO: add order

		for i, ch := range abc {
			if (min == empty || i > order[min]) && count[ch] > 0 && !bl[prev][ch] {
				return ch
			}
		}

		return empty
	}

	i, c := int64(0), uint32(0)
	for i >= 0 && i < l {
		prev, s := empty, empty

		if i > 0 {
			prev = r[i-1]
		}

		if r[i] > empty {
			s = r[i]
		}

		found := false
		next := getOptions(s, prev)

		found = next != empty //&& (i != l-1 || !bl[next][r[0]])

		if found {
			r[i] = next
			count[next]--
			i++
		} else {
			i--
			if i >= 0 && r[i] > empty {
				count[r[i]]++
				if i < l-1 {
					r[i+1] = empty
				}
			}
			c++
		}

	}

	return i > 0, string(r), c
}

func solve(n uint32, s1, s2 []rune) (bool, string) {
	variants := []string{"abc", "bac", "cab", "cba", "bca", "acb"}
	abc := ""

	var m uint32 = 8
	if n < m {
		m = n
	}

	bl := map[rune]map[rune]bool{}

	for _, ch := range [...]rune{'a', 'b', 'c', empty} {
		bl[ch] = map[rune]bool{}
	}

	bl[s1[0]][s1[1]] = true
	bl[s2[0]][s2[1]] = true

	min := uint32(65536)
	for _, v := range variants {
		f, _, c := sample(m, bl, v)
		if f && c < min {
			min, abc = c, v
		}
	}

	if len(abc) == 0 {
		return false, abc
	}

	f, s, _ := sample(n, bl, abc)

	return f, s
}

func main() {
	var n uint32
	var s1, s2 string

	fmt.Scan(&n)
	fmt.Scan(&s1, &s2)

	f, res := solve(n, ([]rune)(s1), ([]rune)(s2))

	if f {
		fmt.Println("YES")
		fmt.Println(res)
	} else {
		fmt.Println("NO")
	}
}
