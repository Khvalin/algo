package main

import (
	"fmt"
	"sort"
)

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

	reduce := func(distances A, cols bool) {
		visitAll(distances, func(r, c, d int) {
			if cols {
				distances[r][c] -= colMins[c]
			} else {
				distances[r][c] -= rowMins[r]
			}

		})
	}

	getZeros := func(distances A) []cell {
		zeros := []cell{}

		visitAll(distances, func(r, c, d int) {
			if d == 0 {
				zeros = append(zeros, cell{r: r, c: c})
			}
		})

		return zeros
	}

	tryAssign := func(cells []cell) ([]int, []int, int, int) {
		type assigment struct {
			count int
			isCol bool
			index int
		}

		assignmentCols, assignmentRows := make([]assigment, L), make([]assigment, L)

		for _, cell := range cells {
			assignmentCols[cell.c].index = cell.c
			assignmentCols[cell.c].isCol = true
			assignmentCols[cell.c].count++

			assignmentRows[cell.r].index = cell.r
			assignmentRows[cell.r].count++
		}

		assignments := []assigment{}
		for i := 0; i < L; i++ {
			if assignmentCols[i].count > 0 {
				assignments = append(assignments, assignmentCols[i])
			}
			if assignmentRows[i].count > 0 {
				assignments = append(assignments, assignmentRows[i])
			}
		}

		sort.Slice(assignments, func(i, j int) bool {
			return assignments[i].count > assignments[j].count
		})
		fmt.Println(assignments)

		n := 0
		for ; len(cells) > 0; n++ {
			ass := assignments[n]
			for j := len(cells) - 1; j >= 0; j-- {
				cell := cells[j]
				if (ass.isCol && ass.index == cell.c) || (!ass.isCol && ass.index == cell.r) {
					cells[j] = cells[len(cells)-1]
					cells = cells[:len(cells)-1]
				}
			}
		}

		cols, rows := make([]int, L), make([]int, L)

		cCount, rCount := 0, 0

		for i := 0; i <= n; i++ {
			ass := assignments[i]
			if ass.isCol {
				cols[ass.index]++
				cCount++
			} else {
				rows[ass.index]++
				rCount++

			}
		}

		return cols, rows, cCount, rCount
	}

	findMin := func(distances A, cols, rows []int) int {
		min := max

		visitAll(distances, func(r, c, d int) {
			if cols[c] == 0 && rows[r] == 0 && d < min {
				min = d
			}
		})

		return min
	}

	subtract := func(distances A, cols, rows []int, n int) {
		visitAll(distances, func(r, c, d int) {
			if cols[c] == 0 && rows[r] == 0 {
				distances[r][c] -= n
			}

			if cols[c] > 0 && rows[r] > 0 {
				distances[r][c]++
			}
		})

	}

	prepare(distances)

	var assignmentCols, assignmentRows []int
	var cCount, rCount int

	for round := 0; cCount+rCount < L; round++ {

		if round < 2 {
			updateMins(distances)
			reduce(distances, round%2 == 0)
		} else {
			min := findMin(distances, assignmentCols, assignmentRows)
			subtract(distances, assignmentCols, assignmentRows, min)
		}

		zeros := getZeros(distances)

		assignmentCols, assignmentRows, cCount, rCount = tryAssign(zeros)
	}

	res := 0

	visitAll(distances, func(r, c, d int) {
		if d == 0 && orig[r][c] > res {
			res = orig[r][c]
		}
	})

	return res
}
