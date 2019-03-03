package main

import "fmt"

func fillShortestPaths(plan [][]string) [][]string {
	findMin := func(ints ...int) int {
		min := ints[0]
		for _, v := range ints {
			if v < min {
				min = v
			}
		}

		return min
	}

	findMax := func(ints ...int) int {
		max := ints[0]
		for _, v := range ints {
			if v > max {
				max = v
			}
		}

		return max
	}

	m, n := len(plan), len(plan[0])
	planWeight := make([][]int, m)

	res := plan
	var start [2]int
	for i := 0; i < m; i++ {
		planWeight[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if plan[i][j] == "s" {
				start = [2]int{i, j}
			}

			minLen := findMin(i, j, m-i-1, n-j-1)
			planWeight[i][j] = minLen
		}
	}

	q := make([][2]int, n*m*2)
	q[0] = start
	j := 1
	i := 0
	for i < j {
		c := q[i]

		minX, minY := findMax(0, c[0]-1), findMax(0, c[1]-1)
		maxX, maxY := findMin(m-1, c[0]+1), findMin(n-1, c[1]+1)

		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				if planWeight[x][y] == planWeight[c[0]][c[1]]-1 {
					res[x][y] = "#"
					//	fmt.Println(j)
					q[j] = [2]int{x, y}
					j++
				}
			}
		}

		i++
	}

	return res
}

func main() {
	input := [][]string{
		{" ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " "},
		{" ", " ", "s", " ", " ", " "},
		{" ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " "}}

	//	fmt.Println(fillShortestPaths(input))

	input = [][]string{
		{" ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " "},
		{" ", " ", "s", " ", " "},
		{" ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " "},
	}
	fmt.Println(fillShortestPaths(input))

	input = [][]string{
		{" ", "s", " "},
		{" ", " ", " "}}

	fmt.Println(fillShortestPaths(input))
}
