package main

import "fmt"

func solve(a, b, c uint64) uint64 {
	res := a
	if b < a {
		res = b
	}

	res += c
	res *= 2

	if a != b {
		res++
	}

	return res
}

func main() {
	var a, b, c uint64

	fmt.Scan(&a, &b, &c)

	res := solve(a, b, c)
	fmt.Println(res)
}
