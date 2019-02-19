package main

import "fmt"

func solve(n int, numbers []int) (bool, [][]int) {
	res := [][]int{}

	return false, res
}

func main() {
	n := 0

	fmt.Scan(&n)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	fmt.Println(solve(n, numbers))
}
