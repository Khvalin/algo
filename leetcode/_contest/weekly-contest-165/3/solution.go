package main

import "fmt"

func countSquares(matrix [][]int) int {
	n, m := len(matrix), len(matrix[0])
	r := 0

	min := n
	if m < min {
		min = m
	}

	for i := 0; i <= min; i++ {

		for x := 0; x < n-i; x++ {
			for y := 0; y < m-i; y++ {
				ok := true

				for k := x; ok && k <= x+i; k++ {
					for l := y; ok && l <= y+i; l++ {
						ok = ok && (matrix[k][l] == 1)
					}
				}

				if ok {
					r++
				}
			}
		}
	}

	return r
}

func main() {
	matrix := [][]int{
		{0, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 1},
	}

	fmt.Println(countSquares(matrix))
}
