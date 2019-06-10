package sudoku

func sudoku(grid [][]int) bool {
	vertical := map[int]int{}
	horizontal := map[int]int{}
	dgts := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, row := range grid {
		m := map[int]bool{}

		for j, d := range row {
			if d > 0 && d < 10 {
				m[d] = true
			}
			vertical[d*10+i/3]++
			horizontal[d*10+j/3]++
		}

		if len(m) != 9 {
			return false
		}
	}

	for _, d := range dgts {
		for i := 0; i < 3; i++ {
			if vertical[d*10+i/3] != 3 || 3 != horizontal[d*10+i/3] {
				return false
			}
		}
	}

	return true
}
