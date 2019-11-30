package main

func removeInterval(intervals [][]int, toBeRemoved []int) [][]int {
	res := [][]int{}

	a, b := toBeRemoved[0], toBeRemoved[1]

	for i := range intervals {
		x, y := intervals[i][0], intervals[i][1]

		if x > a && y < b {
			continue
		}

		if a > x && b < y {
			res = append(res, []int{x, a})
			res = append(res, []int{b, y})
			continue
		}

		if a > x && a < y {
			y = a
		}

		if b > x && b < y {
			x = b
		}

		res = append(res, []int{x, y})
	}

	return res
}
