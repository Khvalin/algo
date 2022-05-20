package main

func uniquePathsWithObstacles(grid [][]int) int {
	for i, row := range grid {
		for j := range row {
			grid[i][j] *= -1
		}
	}

	if grid[0][0] < 0 || grid[len(grid)-1][len(grid[0])-1] < 0 {
		return 0
	}

	for i := range grid {
		for j, v := range grid[i] {
			if v < 0 {
				continue
			}

			left := 0
			if j > 0 && grid[i][j-1] > 0 {
				left = grid[i][j-1]
			}

			top := 0
			if i > 0 && grid[i-1][j] > 0 {
				top = grid[i-1][j]
			}

			if j == 0 && i == 0 {
				v = 1
			}
			grid[i][j] = v + top + left
		}

	}

	return grid[len(grid)-1][len(grid[0])-1]
}
