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

			key2 := dx*nmax + dy

			d := 0
			ps := &s.dist
			if v, f := s.dist[key2]; f {
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
				//s.dist[key2] = d
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
	input := [][]int{{9, 8, 7}, {4, 5, 6}, {3, 2, 1}}
	fmt.Println(longestIncreasingPath(input), 9)

	input = [][]int{{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, {19, 18, 17, 16, 15, 14, 13, 12, 11, 10}, {20, 21, 22, 23, 24, 25, 26, 27, 28, 29}, {39, 38, 37, 36, 35, 34, 33, 32, 31, 30}, {40, 41, 42, 43, 44, 45, 46, 47, 48, 49}, {59, 58, 57, 56, 55, 54, 53, 52, 51, 50}, {60, 61, 62, 63, 64, 65, 66, 67, 68, 69}, {79, 78, 77, 76, 75, 74, 73, 72, 71, 70}, {80, 81, 82, 83, 84, 85, 86, 87, 88, 89}, {99, 98, 97, 96, 95, 94, 93, 92, 91, 90}, {100, 101, 102, 103, 104, 105, 106, 107, 108, 109}, {119, 118, 117, 116, 115, 114, 113, 112, 111, 110}, {120, 121, 122, 123, 124, 125, 126, 127, 128, 129}, {139, 138, 137, 136, 135, 134, 133, 132, 131, 130}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}

	fmt.Println(longestIncreasingPath(input))
}
