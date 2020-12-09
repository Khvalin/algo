package main

import (
	"fmt"
	"sort"
)

func getSkyline(buildings [][]int) [][]int {
	res := [][]int{}

	/*
	   First, register all building start/end points.
	   To distinguish them, we'll register the starts with
	   positive heights and the ends with negative heights.
	*/
	points := [][]int{}
	for _, b := range buildings {
		points = append(points, []int{b[0], b[2]}, []int{b[1], -b[2]})
	}

	/*
	   Then, we sort the points by from leftmost to rightmost
	   x-coordinate. If there's a tie, then we pick first
	   those points with higher height
	*/
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}

		return points[i][0] < points[j][0]
	})

	/*
	   Next, we want to iterate over the points (from leftmost
	   to rightmost, since we sorted them that way) and keep
	   track of the heights we see. If it is positive, we add it;
	   if it is negative, we remove it.

	   We want to keep these heights sorted. This way, at any
	   given coordinate we will know which is the biggest height,
	   which is what we need to calculate our output.

	   Since we want to keep heights sorted, whenever we want to
	   add or remove a height we'll have to use binary search to
	   be efficient.
	*/

	heights := []int{}
	addHeight := func(h int) {
		l, r := 0, len(heights)-1

		for l <= r {
			mid := (l + r) >> 1

			if heights[mid] >= h {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

		if l > len(heights)-1 {
			heights = append(heights, h)
			return
		}

		heights = append(heights[:l+1], heights[l:]...)
		heights[l] = h
	}

	removeHeight := func(h int) {
		l, r := 0, len(heights)-1

		for l <= r {
			mid := (l + r) >> 1

			if heights[mid] >= h {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

		if l < 0 || l > len(heights)-1 {
			return
		}
		heights = append(heights[:l], heights[l+1:]...)
	}

	prevHeight := 0
	for _, p := range points {
		x, h := p[0], p[1]
		if h > 0 {
			addHeight(h)
		} else {
			removeHeight(-h)
		}

		curHeight := 0
		if len(heights) > 0 {
			curHeight = heights[len(heights)-1]
		}

		if curHeight != prevHeight {
			res = append(res, []int{x, curHeight})
			prevHeight = curHeight
		}
	}

	return res
}

func main() {
	//	buildings := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
	buildings := [][]int{{0, 2, 3}, {2, 5, 3}}
	fmt.Println(getSkyline(buildings))
}
