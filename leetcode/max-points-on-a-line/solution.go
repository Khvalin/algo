package main

import (
	"fmt"
	"math"
	"sort"
)

// Point Definition for a point.
type Point struct {
	X int
	Y int
}

func maxPoints(points []Point) int {
	const eps = 0.00001

	calcLine := func(p1 Point, p2 Point) (float32, float32) {
		x1 := float32(p1.X)
		x2 := float32(p2.X)
		y1 := float32(p1.Y)
		y2 := float32(p2.Y)

		k := (y2 - y1) / (x2 - x1)
		b := y1 - k*x1

		return k, b
	}

	sort.Slice(points, func(i, j int) bool {
		p1, p2 := points[i], points[j]
		d1, d2 := p1.X^2+p1.Y^2, p2.X^2+p2.Y^2

		if d1 == d2 {
			return p1.X < p2.X
		}

		return d1 < d2
	})

	counts := []int{}
	newPoints := []Point{}
	for i := 0; i < len(points); i++ {
		counts = append(counts, 1)
		newPoints = append(newPoints, points[i])
		for ; i+1 < len(points) && points[i+1] == points[i]; i++ {
			counts[len(counts)-1]++
		}
	}
	points = newPoints

	result := 0
	if len(points) < 3 {
		for i := 0; i < len(points); i++ {
			result += counts[i]
		}
		return result
	}

	for i := 0; i < len(points)-2; i++ {

		for j := i + 1; j < len(points)-1; j++ {
			count := counts[i] + counts[j]
			k1, b1 := calcLine(points[i], points[j])
			//fmt.Println(points[i], points[j], k1, b1)

			for l := j + 1; l < len(points); l++ {
				if points[i].X == points[l].X && points[i].X == points[j].X {
					count += counts[l]
					continue
				}

				k2, b2 := calcLine(points[j], points[l])

				if math.Abs(float64(k2-k1)) < eps && math.Abs(float64(b2-b1)) < eps {
					count += counts[l]
				}

				//				fmt.Println(points[i], points[j], points[l], k2, b2, count)
			}

			if count > result {
				result = count
			}
		}
	}
	return result
}

func main() {
	/*	test1 := maxPoints([]Point{
			*&Point{1, 1},
			*&Point{3, 2},
			*&Point{5, 3},
			*&Point{4, 1},
			*&Point{2, 3},
			*&Point{1, 4},
		})

		test2 := maxPoints([]Point{
			*&Point{1, 1},
			*&Point{1, 1},
			*&Point{1, 1},
			*&Point{1, 3},
			*&Point{2, 2},
		})

		//[[0,0],[1,1],[0,0]]

			fmt.Println(test1, test2)
			fmt.Println(maxPoints([]Point{}))
			fmt.Println(maxPoints([]Point{*&Point{1, 1}, *&Point{7, 4}}))*/

	//[[0,-1],[0,3],[0,-4],[0,-2],[0,-4],[0,0],[0,0],[0,1],[0,-2],[0,4]]

	fmt.Println(maxPoints([]Point{*&Point{0, 0}, *&Point{1, 1}, *&Point{0, 0}}))

	fmt.Println(maxPoints([]Point{*&Point{0, -1}, *&Point{0, 3}, *&Point{0, -4}, *&Point{0, -2}, *&Point{0, -4}, *&Point{0, 0}, *&Point{0, 0}, *&Point{0, 1}, *&Point{0, -2}, *&Point{0, 4}}))

	fmt.Println(maxPoints([]Point{*&Point{1, 1}, *&Point{10, 10}, *&Point{10, 10},
		*&Point{0, 0}, *&Point{1, 1}, *&Point{0, 0}, *&Point{2, 3}, *&Point{2, 3}, *&Point{2, 3}}))

	fmt.Println(maxPoints([]Point{*&Point{3, 1}, *&Point{12, 3}, *&Point{-6, -1}}))
}
