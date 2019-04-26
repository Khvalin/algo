package main

import (
	"fmt"
)

func solve(a []uint8, f [10]uint8) []uint8 {
	started := false
	for i, cur := range a {
		if f[cur] < cur {
			if started {
				break
			}
			continue
		}

		if f[cur] > cur {
			a[i] = f[cur]
			started = true
		}
	}
	return a
}

func main() {
	var n int

	fmt.Scanln(&n)

	a := make([]uint8, n)
	s := ""
	fmt.Scanln(&s)
	for i := range a {
		a[i] = s[i] - '0'
	}

	f := [10]uint8{}
	for i := 1; i < 10; i++ {
		fmt.Scan(&f[i])
	}

	ans := solve(a, f)
	for _, cur := range ans {
		fmt.Print(cur)
	}
	fmt.Println()
}
