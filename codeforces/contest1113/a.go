package main

import "fmt"

func solve(n, v int) int {
	res := v
	if n-1 <= v {
		res = n - 1
	}
	for i := 2; i < n; i++ {
		if i <= n-v {
			res += i
		}
	}
	return res
}

func main() {
	n, v := 0, 0

	fmt.Scan(&n, &v)
	fmt.Println(solve(n, v))
}
