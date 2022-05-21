package main

func coinChange(coins []int, amount int) int {
	memo := make([]int, amount+1)
	for i := range memo {
		memo[i] = -1
	}
	memo[0] = 0

	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i < c || memo[i-c] < 0 {
				continue
			}

			res := memo[i-c] + 1
			if memo[i] < 0 || memo[i] > res {
				memo[i] = res
			}
		}
	}

	return memo[amount]
}
