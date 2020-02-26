package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	var res []string

	type state struct {
		l, r int
		str  []rune
	}

	q := []state{state{n, n, []rune{}}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		if p.r == 0 && p.l == 0 {
			res = append(res, string(p.str))
			continue
		}

		if p.r < p.l || p.l < 0 {
			continue
		}

		nextl := state{p.l - 1, p.r, append([]rune{}, p.str...)}
		nextl.str = append(nextl.str, '(')

		nextr := state{p.l, p.r - 1, append([]rune{}, p.str...)}
		nextr.str = append(nextr.str, ')')

		q = append(q, nextl, nextr)
	}

	return res
}

func main() {
	fmt.Println(generateParenthesis(5))
}
