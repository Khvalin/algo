package main

import (
	"fmt"
)

func meanScores(S []int) []int {
	s := 0

	for i, c := range S {
		s += c

		S[i] = s / (1 + i)
	}
	return S
}

func main() {
	fmt.Println(meanScores([]int{100, 20, 50, 70, 45}))
}
