package main

import (
	"fmt"
)

const max = 1 << 31

type A [][]int
type visitFn func(int, int, int)

func homeworkGroup(distances A) int {
	L := len(distances)
	rowMins := make([]int, L)
	colMins := make([]int, L)

	visitAll := func(distances A, visit visitFn) {
		for i := range distances { // rows
			for j := range distances[i] { //cols
				visit(i, j, distances[i][j])
			}
		}
	}

	prepare := func(distances A) {
		visitAll(distances, func(r, c, d int) {
			rowMins[r] = max
			colMins[r] = max

			if d == -1 {
				distances[r][c] = max
			}
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

	reduceCols := func(distances A) int {
		rowsMap := map[int]bool{}

		visitAll(distances, func(r, c, d int) {

			if colMins[0] != 0 {
				distances[r][c] -= colMins[c]
			}
			if distances[r][c] == 0 {
				rowsMap[r] = true
			}

		})
		return len(rowsMap)
	}

	reduceRows := func(distances A) int {
		colsMap := map[int]bool{}

		visitAll(distances, func(r, c, d int) {

			if rowMins[0] != 0 {
				distances[r][c] -= rowMins[c]
			}
			if distances[r][c] == 0 {
				colsMap[c] = true
			}

		})
		return len(colsMap)
	}

	prepare(distances)

	zeroRows, zeroCols := 0, 0
	for round := 0; zeroRows < L && zeroCols < L; round++ {
		updateMins(distances)
		if round%2 == 0 {
			zeroRows = reduceCols(distances)
		} else {
			zeroCols = reduceRows(distances)
		}
		fmt.Println(rowMins)
		fmt.Println(colMins)
	}

	return 0
}
