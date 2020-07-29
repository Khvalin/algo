package main

import "fmt"

func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 1
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}

		return b
	}

	dp := make([][]int, len(dungeon))
	for i := range dp {
		dp[i] = make([]int, len(dungeon[i]))
		for j := range dp[i] {
			dp[i][j] = 0
		}
	}

	m := len(dungeon)
	n := len(dungeon[0])

	//fill the queen bottomcorner
	dp[m-1][n-1] = max(1-dungeon[m-1][n-1], 1)

	for i := m - 2; i >= 0; i-- {
		dp[i][n-1] = max(dp[i+1][n-1]-dungeon[i][n-1], 1)
	}

	for i := n - 2; i >= 0; i-- {
		dp[m-1][i] = max(dp[m-1][i+1]-dungeon[m-1][i], 1)
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			down := max(dp[i+1][j]-dungeon[i][j], 1)
			right := max(dp[i][j+1]-dungeon[i][j], 1)
			dp[i][j] = min(down, right)
		}
	}

	return dp[0][0]
}

func main() {
	dungeon := [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}
	//	dungeon = [][]int{{0, 0, 0}, {1, 1, -1}}
	dungeon = [][]int{{1, -3, 3}, {0, -2, 0}, {-3, -3, -3}}
	fmt.Println(calculateMinimumHP(dungeon))
}
