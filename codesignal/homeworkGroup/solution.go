package main

import "fmt"

const max = 1 << 31

type A [][]int
type visitFn func(int, int, int)

type cell struct {
	r, c int
}

func homeworkGroup(distances A) int {
	L := len(distances)
	rowMins := make([]int, L)
	colMins := make([]int, L)
	orig := make([][]int, L)
	for i := range distances {
		orig[i] = make([]int, L)
		copy(orig[i], distances[i])
	}

	visitAll := func(distances A, visit visitFn) {
		for i := range distances { // rows
			for j := range distances[i] { //cols
				if i != j {
					visit(i, j, distances[i][j])
				}
			}
		}
	}

	prepare := func(distances A) {
		visitAll(distances, func(r, c, d int) {
			rowMins[r] = max
			colMins[r] = max
			/*
				if -1 == d {
					distances[r][c] = max
				}*/
		})

	}

	updateMins := func(distances A) {
		visitAll(distances, func(r, c, d int) {
			if d < rowMins[r] {
				rowMins[r] = d
			}

			if d < colMins[c] {
				colMins[c] = d
			}
		})
	}

	reduce := func(distances A, cols bool) []cell {
		zeros := []cell{}

		visitAll(distances, func(r, c, d int) {
			if cols {
				distances[r][c] -= colMins[c]
			} else {
				distances[r][c] -= rowMins[r]
			}
			if distances[r][c] == 0 {
				zeros = append(zeros, cell{r: r, c: c})
			}

		})
		return zeros
	}

	prepare(distances)

	zeroMap := make(map[cell]bool)
	for round := 0; round < 2; round++ {
		updateMins(distances)
		zeros := reduce(distances, round%2 == 0)

		for _, z := range zeros {
			zeroMap[z] = true
		}
	}
	fmt.Println(zeroMap)
	fmt.Println(distances)

	res := 0

	visitAll(distances, func(r, c, d int) {
		if d == 0 && orig[r][c] > res {
			res = orig[r][c]
		}
	})

	return res
}
