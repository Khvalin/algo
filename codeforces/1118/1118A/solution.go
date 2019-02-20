package main

import "fmt"

func solve(n, a, b uint64) uint64 {
	res := uint64(0)

	if n%2 == 1 {
		res += a
		n--
	}

	if n > 0 {
		if 2*a <= b {
			res += n * a
		} else {
			res += b * (n / 2)
		}
	}

	return res
}

func main() {
	q := 0
	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		var n, a, b uint64
		fmt.Scan(&n, &a, &b)
		fmt.Println(solve(n, a, b))
	}
}
