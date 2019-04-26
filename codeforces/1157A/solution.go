package main

import (
	"fmt"
)

func solve(n int) int {
	m := map[int]bool{}
	m[n] = true

	for ok := false; !ok; {

		n++

		for n > 0 && n%10 == 0 {
			n /= 10
		}

		ok = m[n]
		m[n] = true
	}

	return len(m)
}

func main() {
	var n int

	fmt.Scan(&n)

	ans := solve(n)
	fmt.Println(ans)
}
