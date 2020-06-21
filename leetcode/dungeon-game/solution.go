package main

import "fmt"

func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 1
	}

	minHP := make([][]int, len(dungeon))
	for i := range minHP {
		minHP[i] = make([]int, len(dungeon[i]))
		for j := range minHP[i] {
			minHP[i][j] = 0
		}
	}

	res := 1
	q := [][2]int{}

	x, y := 0, 0
	for x < len(dungeon) || y < len(dungeon[0]) {
		if y < len(dungeon[0]) {
			for i := x; i < len(dungeon); i++ {
				q = append(q, [2]int{i, y})
			}
		}

		if x < len(dungeon) {
			for j := y + 1; j < len(dungeon[x]); j++ {
				q = append(q, [2]int{x, j})
			}
		}
		x++
		y++
	}

	for len(q) > 0 {
		v := q[0]
		q = q[1:]

		x, y, m, s := v[0], v[1], 1, 0
		fmt.Println(x, y)
		if x > 0 {
			m = minHP[x-1][y]
			s = dungeon[x-1][y]
		}
		if y > 0 && (minHP[x][y-1] > m) {
			m = minHP[x][y-1]
			s = dungeon[x][y-1]
		}

		s += dungeon[x][y]
		if s < 0 && s < m {
			m = s
		}
		minHP[x][y] = m
		dungeon[x][y] = s

		if x == len(dungeon)-1 && y == len(dungeon[x])-1 {
			res = m
		}
	}

	fmt.Println(minHP)

	return res
}

func main() {
	dungeon := [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}
	fmt.Println(calculateMinimumHP(dungeon))
}
