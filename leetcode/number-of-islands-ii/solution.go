package main

import "fmt"

// Definition for a point.
type Point struct {
	X, Y int
}

/**
 * @param n: An integer
 * @param m: An integer
 * @param operators: an array of point
 * @return: an integer array
 */
func numIslands2(n int, m int, operators []*Point) []int {
	// write your code here

	l := len(operators)
	points := map[Point]int{}
	components := make([]int, l)

	c := 0
	res := make([]int, l+1)
	//res[l] = -1

	for i, op := range operators {
		if _, found := points[*op]; !found {
			x, y := op.X, op.Y
			pts := []Point{Point{x - 1, y}, Point{x + 1, y}, Point{x, y - 1}, Point{x, y + 1}}

			cIndex := i

			for _, p := range pts {
				if ind, f := points[p]; f {
					if v := components[ind]; v != cIndex {
						cIndex = v
					}
				}
			}

			c++
			points[*op] = i
			components[i] = i

			pts = append(pts, *op) // !!!!

			unique := map[int]bool{}
			for _, p := range pts {
				if ind, f := points[p]; f && components[ind] != cIndex {
					v := components[ind]
					if !unique[v] {
						unique[v] = true
						c--
					}
					components[ind] = cIndex
				}
			}
		}

		res[i] = c
	}

	return res
}

func main() {
	n, m := 8, 14
	operators := []*Point{}

	// operators = []*Point{&Point{0, 1}, &Point{1, 0}, &Point{2, 1}, &Point{1, 2}, &Point{1, 1}}
	// operators = []*Point{&Point{0, 1}, &Point{1, 2}, &Point{2, 1}, &Point{1, 0}, &Point{0, 2}, &Point{0, 0}, &Point{1, 1}}
	operators = []*Point{&Point{0, 9}, &Point{5, 4}, &Point{0, 12}, &Point{6, 9}, &Point{6, 5}, &Point{0, 4}, &Point{4, 11}, &Point{0, 0}, &Point{3, 5}, &Point{6, 7}, &Point{3, 12}, &Point{0, 5}, &Point{6, 13}, &Point{7, 5}, &Point{3, 6}, &Point{4, 4}, &Point{0, 8}, &Point{3, 1}, &Point{4, 6}, &Point{6, 1}, &Point{5, 12}, &Point{3, 8}, &Point{7, 0}, &Point{2, 9}, &Point{1, 4}, &Point{3, 0}, &Point{1, 13}, &Point{2, 13}, &Point{6, 0}, &Point{6, 4}, &Point{0, 13}, &Point{0, 3}, &Point{7, 4}, &Point{1, 8}, &Point{5, 5}, &Point{5, 7}, &Point{5, 10}, &Point{5, 3}, &Point{6, 10}, &Point{6, 2}, &Point{3, 10}, &Point{2, 7}, &Point{1, 12}, &Point{5, 0}, &Point{4, 5}, &Point{7, 13}, &Point{3, 2}}
	fmt.Println(numIslands2(n, m, operators))
}
