package main

import "fmt"

type state struct {
	x, y int
	dist map[int]int
}

func longestIncreasingPath(matrix [][]int) int {
	nmax := 201
	st := &(map[int]int{})

	var solve func(s state) int
	solve = func(s state) int {
		res := 1
		x, y := s.x, s.y

		key := nmax*x + y

		if v, f := s.dist[key]; f {
			return v
		}

		//s.dist[key] = 0

		dirs := [][2]int{{x + 1, y}, {x - 1, y}, {x, y - 1}, {x, y + 1}}

		for _, dir := range dirs {
			dx, dy := dir[0], dir[1]

			if dx < 0 || dy < 0 || dx >= len(matrix) || dy >= len(matrix[0]) || matrix[dx][dy] <= matrix[x][y] {
				continue
			}

			key := dx*nmax + dy

			d := 0
			ps := &s.dist
			if v, f := s.dist[key]; f {
				d = v + 1
			} else {
				next := state{dx, dy, map[int]int{}}
				for k, v := range s.dist {
					next.dist[k] = v
				}
				ps = &next.dist

				d = solve(next) + 1
			}

			if d > res {
				res = d
				s.dist[key] = d
				st = ps
			}
		}

		s.dist[key] = res

		return res
	}

	res := 0
	for x, row := range matrix {
		for y := range row {
			r := solve(state{x, y, *st})
			if r > res {
				res = r
			}
		}
	}

	return res
}

func main() {
	input := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}

	fmt.Println(longestIncreasingPath(input))
}
