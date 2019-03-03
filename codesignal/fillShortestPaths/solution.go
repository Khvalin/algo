package main

import "fmt"

func fillShortestPaths(plan [][]string) [][]string {

	type coord struct {
		x, y int
	}

	res := plan
	start := coord{}
	for i := 0; i < len(plan); i++ {
		for j := 0; j < len(plan[i]); j++ {
			if plan[i][j] == "s" {
				start.x, start.y = i, j
			}
		}
	}

	q := []coord

	fmt.Println(start)

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

	fmt.Println(fillShortestPaths(input))
}
