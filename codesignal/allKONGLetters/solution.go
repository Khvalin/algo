package main

import "fmt"

import (
	"math"
	"strings"
)

const wallChar = '#'
const startChar = "S"
const endChar = 'E'
const fishChar = 'F'

// Point is here
type Point struct {
	x, y  int
	fish  bool
	flags byte
}

func allKONGLetters(stage []string) int {
	processInput := func(stage *[]string) (Point, Point) {
		var start Point
		var end Point

		stg := *stage
		//top := strings.Repeat(string(wallChar), len(stg[0])+2)

		for i := 0; i < len(stg); i++ {
			//  stg[i] = string(wallChar) + stg[i] + string(wallChar)
			sIndex := strings.Index(stg[i], startChar)
			eIndex := strings.IndexRune(stg[i], endChar)

			if sIndex > -1 {
				start.x, start.y = i, sIndex
			}

			if eIndex > -1 {
				end.x, end.y = i, eIndex
			}
		}

		end.flags = 1 | 2 | 3 | 4 | 8
		//*stage = append([]string{top}, *stage...)
		//*stage = append(*stage, top)

		return start, end
	}

	if 0 == len(stage) || 0 == len(stage[0]) {
		return -1
	}

	chars := (map[byte]byte{
		'K': 1,
		'O': 2,
		'N': 4,
		'G': 8,
	})

	start, end := processInput(&stage)
	//	fmt.Println(start, end)

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
		push := func(x, y int, flags byte) {
			newP := Point{x, y, fish, flags}
			vectors = append(vectors, newP)
		}

		for k := -1; k < 2; k += 2 {
			flags := cur.flags
			for x := 1; x <= maxX; x++ {
				dx := cur.x + x*k
				if dx >= 0 && dx < len(stage) && wallChar != stage[dx][cur.y] &&
					(endChar != stage[dx][cur.y] || flags == end.flags) {
					flags |= chars[stage[dx][cur.y]]
					push(dx, cur.y, flags)
				} else {
					break
				}
			}

			flags = cur.flags
			for y := 1; y <= maxY; y++ {
				dy := cur.y + y*k
				if dy >= 0 && dy < len(stage[0]) && wallChar != stage[cur.x][dy] &&
					(endChar != stage[cur.x][dy] || flags == end.flags) {
					flags |= chars[stage[cur.x][dy]]
					push(cur.x, dy, flags)
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

	end.flags = 1 | 2 | 4 | 8

	for i := 0; i < 2; i++ {
		end.fish = i == 1
		d, found := distance[end]

		if found && d < result {
			result = d
		}
	}

	/*
		// test output
		for i := 0; i < len(stage); i++ {
			for j := 0; j < len(stage[i]); j++ {
				dist, _ := distance[Point{i, j, true, 1 | 2 | 4 | 8}]
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
	fmt.Println(allKONGLetters([]string{
		"########",
		"# K O###",
		"#F######",
		"#SN G  E",
		"########",
	}), 7)

	/*
		fmt.Println(allKONGLetters([]string{
			"########",
			"# KONG #",
			"# #### #",
			"#S    E#",
			"########",
		}))
	*/
}
