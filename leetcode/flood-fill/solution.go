package main

import (
	"fmt"
)

type point struct {
	X int
	Y int
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}

	stack := []point{point{X: sr, Y: sc}}
	startColor := image[sr][sc]
	candidates := [4]point{}

	for len(stack) > 0 {
		var p point
		p, stack = stack[0], stack[1:]
		image[p.X][p.Y] = newColor
		//fmt.Println(p, image, stack)

		candidates[0].X, candidates[0].Y = p.X-1, p.Y
		candidates[1].X, candidates[1].Y = p.X+1, p.Y
		candidates[2].X, candidates[2].Y = p.X, p.Y+1
		candidates[3].X, candidates[3].Y = p.X, p.Y-1

		for i := 0; i < len(candidates); i++ {
			if candidates[i].X >= 0 && candidates[i].Y >= 0 && candidates[i].X < len(image) && candidates[i].Y < len(image[0]) && startColor == image[candidates[i].X][candidates[i].Y] {
				stack = append(stack, candidates[i])
			}
		}
		//	fmt.Println(stack)
	}

	return image
}

func main() {
	fmt.Println(floodFill([][]int{[]int{1, 1, 1}, []int{1, 1, 0}, []int{1, 0, 1}}, 1, 1, 2))
}
