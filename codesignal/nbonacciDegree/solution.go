package main

import (
	"fmt"
)

func nbonacciDegree(sequence []int) int {
	sums := make([]int, len(sequence)+1)
	degrees := map[int]bool{}

	total := 0
	res := -1
	for i, n := range sequence {
		for k := range degrees {
			sums[k] += sequence[i-1] - sequence[i-k-1]
			if sums[k] != n {
				if k <= res {
					res = -1
				}
				delete(degrees, k)
			} else {
				if k > res {
					res = k
				}
			}
		}

		sums[i] = total
		if total == n && i > 0 {
			degrees[i] = true
			res = i
		}
		total += n
	}

	//	fmt.Println(degrees)

	return res
}

func main() {
	fmt.Println(nbonacciDegree([]int{1, 2, 3}), 2)
	fmt.Println(nbonacciDegree([]int{0, 0, 0, 0, 0, 0, 0}), "??")
	fmt.Println(nbonacciDegree([]int{-1, 2, 0, 0, 1, 1, 2, 1, 1, -1, -1, 0, 1, -1, 0, -1, -2, 0, 0, 1, -1, 1, 0, 1, 2, -2, 1, 0, 0, 1, 1, -1, 1, 1, 2, 1, 2, 2, 0, -2, 0, -2, 0, -1, 1, 1, -2, 2, 0, -2, -2, -1, -2, -1, 2, 1, 2, -1, -2, 1, 0, 2, -1, 1, -1, -2, -1, -1, 0, 0, 0, 2, -2, -2, -2, 2, 2, -2, -2, 1, 1, -2, 0, -2, -1, -2, -1, 0, -1, -2, 2, 2, 2, 2, 1, 1, -1, -1, -2, 0, -3, -5, -12, -24, -48, -97, -195, -392, -785, -1571, -3141, -6281, -12562, -25125, -50249, -100498, -200995, -401988, -803976, -1607952, -3215905, -6431809, -12863619, -25727238, -51454477, -102908956}), 100)

	fmt.Println(nbonacciDegree([]int{0, 6, -2, 3, 7, 14, 22, 46, 89, 171, 328, 634, 1222, 2355}), 4)
}
