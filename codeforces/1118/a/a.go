package main

import "fmt"

func solve(n, a, b int) int {
	res := 0
	if a*2 >= b {
		res += b * (n / 2)
		n = n % 2
	}
	res += n * a

	return res
}

func main() {
	q, n, a, b := 0, 0, 0, 0

	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		fmt.Scan(&n, &a, &b)
		fmt.Println(solve(n, a, b))
	}
}
