package main

import "fmt"

const max = 1 << 31

type A [][]int
type visitFn func(int, int, int)

type cell struct {
	r, c int
}

func (c cell) String() string {
	return fmt.Sprintf("{r: %v, c: %v}", c.r, c.c)
}

type assigment struct {
	count int
	isCol bool
	index int
}

func (a assigment) String() string {
	dir := "row"
	if a.isCol {
		dir = "col"
	}
	return fmt.Sprintf("{ %v: %v, %v }", dir, a.index, a.count)
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

	tryAssign := func(zeros []cell) ([]bool, []bool, int, int) {
		cols, rows, zeroRows := make([]bool, L), make([]bool, L), make([]bool, L)
		zeroRowsMap := make([]int, L)

		starred := make([]bool, L)
		colsCount := 0
		for i, z := range zeros {
			if !cols[z.c] && !zeroRows[z.r] {
				starred[i] = true
				cols[z.c] = true
				zeroRows[z.c] = true
				zeroRowsMap[z.c] = i

				colsCount++
			}
		}

		// Terminate the algorithm if all columns are covered.
		if colsCount < L {
			/*
					Step 3:
				Main Zero Search

				Find an uncovered Z in the distance matrix and prime it Z -> Z' . If no such zero exists, go to Step 5
				If No Z* exists in the row of the Z', go to Step 4.
				If a Z* exists, cover this row and uncover the column of the Z*. Return to Step 3.1 to find a new Z.
			*/
			for i, z := range zeros {
				if !starred[i] {
					if !zeroRows[z.r] {
						break //TODO:??
					}
					rows[z.r] = true
					cols[zeroRowsMap[z.r]] = false
				}
			}
		}
		cCount, rCount := 0, 0

		for _, ass := range finalAssignments {
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

	findMin := func(distances A, cols, rows []bool) int {
		min := max

		visitAll(distances, func(r, c, d int) {
			if !cols[c] && !rows[r] && d < min {
				min = d
			}
		})

		return min
	}

	subtract := func(distances A, cols, rows []bool, n int) {
		visitAll(distances, func(r, c, d int) {
			if !rows[r] && !cols[c] {
				distances[r][c] -= n
			}

			if cols[c] && rows[r] {
				distances[r][c] += n
			}
		})

	}

	prepare(distances)

	var assignmentCols, assignmentRows []bool
	var cCount, rCount int

	for round := 0; cCount+rCount < L; round++ {
		if round < 2 {
			updateMins(distances)
			reduce(distances, round%2 == 0)
		} else {
			min := findMin(distances, assignmentCols, assignmentRows)
			fmt.Println(min, assignmentCols, assignmentRows)
			subtract(distances, assignmentCols, assignmentRows, min)
		}

		zeros := getZeros(distances)

		assignmentCols, assignmentRows, cCount, rCount = tryAssign(zeros)
		fmt.Println(cCount, rCount)
	}

	res := 0

	visitAll(distances, func(r, c, d int) {
		if d == 0 && orig[r][c] > res {
			res = orig[r][c]
		}
	})

	return res
}
