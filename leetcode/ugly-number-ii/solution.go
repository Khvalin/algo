package main

import "fmt"

func nthUglyNumber(n int) int {
	min := func(nums ...int) int {
		res := nums[0]
		for _, n := range nums {
			if n < res {
				res = n
			}
		}

		return res
	}

	two, three, five := 0, 0, 0

	uglies := make([]int, n+1)
	uglies[0] = 1

	for i := 1; i < n; i++ {
		uglies[i] = min(uglies[five]*5, uglies[three]*3, uglies[two]*2)
		if uglies[i] == uglies[two]*2 {
			two++
		}

		if uglies[i] == uglies[three]*3 {
			three++
		}

		if uglies[i] == uglies[five]*5 {
			five++
		}
	}

	return uglies[n-1]
}

func main() {
	fmt.Println(nthUglyNumber(10))
}
