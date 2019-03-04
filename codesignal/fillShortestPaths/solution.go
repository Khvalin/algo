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

	visited := map[[2]int]bool{}

	q := make([][2]int, (n+1)*(m+1))
	//	fmt.Println(len(q))
	q[0] = start
	visited[start] = true

	j := 1
	i := 0
	for i < j {
		c := q[i]

		minX, minY := findMax(0, c[0]-1), findMax(0, c[1]-1)
		maxX, maxY := findMin(m-1, c[0]+1), findMin(n-1, c[1]+1)

		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				cur := [2]int{x, y}
				if !visited[cur] && planWeight[x][y] == planWeight[c[0]][c[1]]-1 {
					visited[cur] = true

					res[x][y] = "#"
					q[j][0], q[j][1] = cur[0], cur[1]
					//fmt.Println(j, x, y, planWeight[x][y], planWeight[c[0]][c[1]])
					j++
				}
			}
		}

		i++
	}

	return res
}

func main() {
	/* 	input := [][]string{
	   		{" ", " ", " ", " ", " ", " "},
	   		{" ", " ", " ", " ", " ", " "},
	   		{"s", " ", " ", " ", " ", " "},
	   		{" ", " ", " ", " ", " ", " "},
	   		{" ", " ", " ", " ", " ", " "},
	   		{" ", " ", " ", " ", " ", " "},
	   		{" ", " ", " ", " ", " ", " "},
	   		{" ", " ", " ", " ", " ", " "}}

	   	fmt.Println(fillShortestPaths(input))
	*/
	input2 := make([][]string, 99)
	for i := range input2 {
		input2[i] = make([]string, 99)

		for k := 0; k < 99; k++ {
			input2[i][k] = " "
		}
	}

	input2[45][45] = "s"

	fmt.Println(fillShortestPaths(input2))

	/*
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

		fmt.Println(fillShortestPaths(input)) */
}
