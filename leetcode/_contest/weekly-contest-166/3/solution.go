package main

import (
	"fmt"
	"sort"
)

func smallestDivisor(nums []int, threshold int) int {
	m := nums[0]
	for _, n := range nums {
		if n > m {
			m = n
		}
	}

	r := sort.Search(m-1, func(i int) bool {
		i = i + 1

		s := 0
		for _, n := range nums {
			s += n / i
			if n%i > 0 {
				s++
			}
		}

		return s <= threshold
	})

	return r + 1
}

func main() {
	fmt.Println(5, smallestDivisor([]int{1, 2, 5, 9}, 6))
	fmt.Println(3, smallestDivisor([]int{2, 3, 5, 7, 11}, 11))
	fmt.Println(4, smallestDivisor([]int{19}, 5))
}
