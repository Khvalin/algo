package main

//package main

import "fmt"

func tallestBillboard(rods []int) int {
	findMax := func(nums ...int) int {
		res := nums[0]
		for _, n := range nums {
			if n > res {
				res = n
			}
		}

		return res
	}

	N := len(rods)
	dp := [21][10001]bool{}
	max := [21][10001]int{}
	dp[0][5000] = true

	for i := 0; i < N; i++ {
		for j := 0; j <= 10000; j++ {
			if j-rods[i] >= 0 && dp[i][j-rods[i]] {
				dp[i+1][j] = true
				max[i+1][j] = findMax(max[i+1][j], max[i][j-rods[i]]+rods[i])
			}
			if j+rods[i] <= 10000 && dp[i][j+rods[i]] {
				dp[i+1][j] = true
				max[i+1][j] = findMax(max[i+1][j], max[i][j+rods[i]])
			}
			if dp[i][j] {
				dp[i+1][j] = true
				max[i+1][j] = findMax(max[i+1][j], max[i][j])
			}
		}
	}
	return max[N][5000]
}

func main() {

	fmt.Println(tallestBillboard([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(tallestBillboard([]int{1, 2, 3, 6}))
	fmt.Println(tallestBillboard([]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 50, 50, 50, 150, 150, 150, 100, 100, 100, 123}))

	fmt.Println(tallestBillboard([]int{124, 121, 107, 127, 156, 146, 135, 153, 137, 150, 141, 138, 129, 142, 124, 144, 126, 900, 900, 900}))

}
