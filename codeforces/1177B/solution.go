package main

import "fmt"

type int = uint64

// Solve returns k-th char from 123456789101112131415161718192021222324252627282930313233343536
func Solve(n int) int {
	return n
}

func main() {
	var n int

	fmt.Scan(&n)

	ans := Solve(n)
	fmt.Println(ans)
}
