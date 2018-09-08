package main

import (
	"fmt"
	"math"
	"strings"
)

const wallChar = '#'
const startChar = "S"
const endChar = "E"
const fishChar = 'F'

// Point is here
type Point struct {
	x, y int
	fish bool
}

func toolAssistedSpeedrun(stage []string) int {
	processInput := func(stage *[]string) (Point, Point) {
		var start Point
		var end Point

		stg := *stage
		//top := strings.Repeat(string(wallChar), len(stg[0])+2)

		for i := 0; i < len(stg); i++ {
			//  stg[i] = string(wallChar) + stg[i] + string(wallChar)
			sIndex := strings.Index(stg[i], startChar)
			eIndex := strings.Index(stg[i], endChar)

			if sIndex > -1 {
				start.x, start.y = i, sIndex
			}

			if eIndex > -1 {
				end.x, end.y = i, eIndex
			}
		}
		//*stage = append([]string{top}, *stage...)
		//*stage = append(*stage, top)

		return start, end
	}

	if 0 == len(stage) || 0 == len(stage[0]) {
		return -1
	}

	start, end := processInput(&stage)
	fmt.Println(start, end)

	distance := make(map[Point]int)
	distance[start] = 0

	q := []Point{start}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		maxX, maxY, fish := 1, 1, cur.fish || (fishChar == stage[cur.x][cur.y])
		if fish {
			maxX, maxY = 2, 3
		}

		vectors := []Point{}
		for k := -1; k < 2; k += 2 {
			for x := 1; x <= maxX; x++ {
				dx := cur.x + x*k
				if dx >= 0 && dx < len(stage) && wallChar != stage[dx][cur.y] {
					vectors = append(vectors, Point{dx, cur.y, fish})
				} else {
					break
				}
			}

			for y := 1; y <= maxY; y++ {
				dy := cur.y + y*k
				if dy >= 0 && dy < len(stage[0]) && wallChar != stage[cur.x][dy] {
					vectors = append(vectors, Point{cur.x, dy, fish})
				} else {
					break
				}
			}
		}

		for _, next := range vectors {
			dist, found := distance[next]

			if !found || dist > distance[cur]+1 {
				distance[next] = distance[cur] + 1
				q = append(q, next)
			}
		}
	}

	result := math.MaxInt32

	for i := 0; i < 2; i++ {
		end.fish = i%2 == 1
		d, found := distance[end]

		if found && d < result {
			result = d
		}
	}

	/*
		// test output
		for i := 0; i < len(stage); i++ {
			for j := 0; j < len(stage[i]); j++ {
				dist, _ := distance[Point{i, j, true}]
				fmt.Printf("%3d", dist)
			}
			fmt.Println()
		}
		// test output end
	*/
	if result == math.MaxInt32 {
		return -1
	}
	return result
}

func main() {
	fmt.Println(toolAssistedSpeedrun([]string{
		"####################",
		"S                 E#",
		"#                  #",
		"####  ###########  #",
		"#####              #",
		"#####  F           #",
		"####################",
	}))

	fmt.Println(toolAssistedSpeedrun([]string{
		"##########",
		"##########",
		"##########",
		"###F######",
		"### ######",
		"S      E #",
		"# ###### #",
		"#        #",
		"##########",
		"##########",
	}))
}
