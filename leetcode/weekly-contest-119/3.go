package main

import "fmt"

func subarraysDivByK(A []int, K int) int {
	res := 0
	p := make([]int, len(A))

	for i, n := range A {
		if i > 0 {
			p[i] = p[i-1]
		}

		p[i] += n
	}

	fmt.Println(p)

	return res
}

func main() {
	fmt.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
}
