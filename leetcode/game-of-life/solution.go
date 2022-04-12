
func gameOfLife(board [][]int) {
	for i, row := range board {
		for j, n := range row {
			count := -(n & 1)
			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if x >= 0 && y >= 0 && x < len(board) && y < len(row) {
						count += board[x][y] & 1
					}
				}
			}

			next := 0
			if n%2 == 1 {
				if count == 2 || count == 3 {
					next = 1
				}
			}

			if n%2 == 0 && count == 3 {
				next = 1
			}

			board[i][j] = next<<1 + n
		}
	}

	for i, row := range board {
		for j := range row {
			board[i][j] = board[i][j] >> 1
		}
	}
}
