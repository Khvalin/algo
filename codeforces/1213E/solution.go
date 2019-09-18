package main

import (
	"fmt"
)

const empty = '!'

func solve(n uint32, s1, s2 []rune) (bool, string) {
	const ABC = "abc"

	m := uint32(n)
	if n < m {
		m = n
	}
	l := int64(3 * m)

	count := map[rune]uint32{
		'a': m,
		'b': m,
		'c': m,
	}

	r := make([]rune, l)
	bl := map[rune]map[rune]bool{}

	for _, ch := range [...]rune{'a', 'b', 'c', empty} {
		bl[ch] = map[rune]bool{}
	}

	bl[s1[0]][s1[1]] = true
	bl[s2[0]][s2[1]] = true

	getOptions := func(min, prev rune) rune {
		//TODO: add order

		for _, ch := range ABC {
			if ch >= min && count[ch] > 0 && !bl[prev][ch] {
				return ch
			}
		}

		return empty
	}

	i := int64(0)
	for i >= 0 && i < l {
		prev, s := empty, empty

		if i > 0 {
			prev = r[i-1]
		}

		if r[i] > empty {
			s = r[i] + 1
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
		}

		//	fmt.Println(string(r), count)
	}

	if i < 0 {
		return false, ""
	}

	res := string(r)
	/* if n > m {
		res = strings.Repeat(res, int(n>>1))

		if n%2 > 0 {
			r = []rune{empty, empty, empty}

			count['a'] = 1
			count['b'] = 1
			count['c'] = 1

			for i := 0; i < 3 && i >= 0; {
				prev, s := rune(res[len(res)-1]), empty

				if i > 0 {
					prev = r[i-1]
				}

				if r[i] > empty {
					s = r[i] + 1
				}

				found := false
				next := getOptions(s, prev)

				found = next != empty

				if found {
					count[next]--
					r[i] = next
					i++
				} else {
					i--
					if r[i] > empty {
						count[r[i]]++
						if i < len(r)-1 {
							r[i+1] = empty
						}
					}
				}
			}

			res += string(r)
		}
	} */

	return true, res
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
