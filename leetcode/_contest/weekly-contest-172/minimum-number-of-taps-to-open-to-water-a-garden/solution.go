package main

import "fmt"

func minTaps(n int, ranges []int) int {
	nmax := 65536
	memo := map[int]int{}

	var solve func(x, ind int) int
	solve = func(x, ind int) int {
		if x >= n {
			return 0
		}

		if ind >= len(ranges) {
			return nmax
		}

		key := x*10001 + ind
		r, f := memo[key]
		if f {
			return r
		}

		r = nmax
		for i := ind; i < n+1; i++ {
			if i-ranges[i] > x || i+ranges[i] < x {
				continue
			}

			t := 1 + solve(i+ranges[i], i+1)
			if t < r {
				r = t
			}
		}

		memo[key] = r

		return r
	}

	r := solve(0, 0)

	if r < nmax {
		return r
	}

	return -1
}

func main() {
	//n := 5
	ranges := []int{0, 0, 0, 0, 0}
	//fmt.Println(minTaps(n, ranges))

	ranges = []int{4, 0, 0, 0, 0, 0, 0, 0, 4}
	fmt.Println(minTaps(8, ranges))
}
