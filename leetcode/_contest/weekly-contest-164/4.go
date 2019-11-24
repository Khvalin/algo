package main

import "fmt"

// type state struct {
// 	ind   int
// 	steps int
// }

func numWays(steps int, arrLen int) int {
	MOD := 1000000007

	state := make([]map[int]int, arrLen+1)

	var solve func(ind, s int) int
	solve = func(ind, s int) int {
		if s == 0 {
			if ind == 0 {
				return 1
			}

			return 0
		}

		if state[ind] == nil {
			state[ind] = map[int]int{}
		}

		r, f := state[ind][s]
		if f {
			return r
		}

		for i := -1; i <= 1; i++ {
			next := ind + i
			if next >= 0 && next < arrLen {
				r = (r + solve(next, s-1)) % MOD
			}
		}

		state[ind][s] = r

		return r
	}

	r := solve(0, steps)

	return r
}

func main() {
	fmt.Println(4, numWays(3, 2))
	fmt.Println(2, numWays(2, 4))
	fmt.Println(8, numWays(4, 2))
}
