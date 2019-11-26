package main

import "fmt"

type grid = [13][13]bool

func tilingRectangle(n int, m int) int {
	createGrid := func(c *grid) grid {
		res := grid{}
		return res
	}

	var solve func(g grid) int
	solve = func(g grid) int {
		return 0
	}

	g := createGrid(nil)

	return solve(g)
}

func main() {
	fmt.Println("3 == ", tilingRectangle(2, 3))
}
