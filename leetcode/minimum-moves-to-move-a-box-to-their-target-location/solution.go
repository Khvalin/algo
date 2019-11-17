package main

import "fmt"

type point struct {
	x int
	y int
}

func minPushBox(grid [][]byte) int {
	var dfs func(start point) map[point]int

	dfs = func(start point) map[point]int {
		res := map[point]int{start: 0}

		points := []point{start}

		for len(points) > 0 {
			p := points[0]
			d := res[p]
			points = points[1:]

			for dd := -1; dd <= 1; dd += 2 {
				p1, p2 := point{p.x + dd, p.y}, point{p.x, p.y + dd}
				d1, f1 := res[p1]
				d2, f2 := res[p2]

				if p1.x >= 0 && p1.x < len(grid) && grid[p1.x][p1.y] != '#' && (!f1 || d1 > d+1) {
					points = append(points, p1)
					res[p1] = d + 1
				}

				if p2.y >= 0 && p2.y < len(grid[0]) && grid[p2.x][p2.y] != '#' && (!f2 || d2 > d+1) {
					points = append(points, p2)
					res[p2] = d + 1
				}
			}
		}

		return res
	}

	r := -1

	var T, B, S point
	_, _, _ = T, B, S

	for x, r := range grid {
		for y, ch := range r {
			switch ch {
			case 'T':
				T = point{x, y}
			case 'B':
				B = point{x, y}
			case 'S':
				S = point{x, y}
			}
		}
	}

	sMap := (dfs(S))
	_, f := sMap[T]
	if !f {
		return -1
	}
	//	tMap := dfs(T)

	return r
}

func main() {
	grid := [][]byte{
		{'#', '#', '#', '#', '#', '#'},
		{'#', 'T', '#', '#', '#', '#'},
		{'#', '.', '.', 'B', '.', '#'},
		{'#', '.', '#', '#', '.', '#'},
		{'#', '.', '.', '.', 'S', '#'},
		{'#', '#', '#', '#', '#', '#'}}

	fmt.Println(minPushBox(grid))
}
