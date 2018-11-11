package main

import "fmt"

const max = 1 << 31

type A [][]int
type visitFn func(int, int, int) bool
type munkersStep func() int

type cell struct {
	r, c int
}

func (c cell) String() string {
	return fmt.Sprintf("{r: %v, c: %v}", c.r, c.c)
}

func homeworkGroup(distances A) int {
	L := len(distances)
	orig := make([][]int, L)
	for i := range distances {
		orig[i] = make([]int, L)
		copy(orig[i], distances[i])
	}

	visitAll := func(distances A, visit visitFn) {
		for i := range distances { // rows
			for j := range distances[i] { //cols
				//	if i != j {
				if !visit(i, j, distances[i][j]) {
					return
				}
				//}
			}
		}
	}

	prepare := func(distances A) {
		visitAll(distances, func(r, c, d int) bool {
			if -1 == d {
				distances[r][c] = max
			}
			return true
		})
	}

	munkers := func() []cell {
		starredZeros, primedZeros := []cell{}, []cell{}
		coveredCols, coveredRows := make([]bool, L), make([]bool, L)

		printMatrix := func() { /*
				for _, r := range distances {
					fmt.Println(r)
				}
				fmt.Print(starredZeros, "\n\n\n")
			*/
		}

		step1 := func() int {
			/*
				Step 1:  For each row of the matrix,
				find the smallest element and subtract it from every element in its row.
				Go to Step 2.
			*/

			rowMins := make([]int, L)
			for r := range rowMins {
				rowMins[r] = max
			}

			visitAll(distances, func(r, c, d int) bool {
				if d < rowMins[r] {
					rowMins[r] = d
				}
				return true
			})

			fmt.Println(rowMins)

			visitAll(distances, func(r, c, d int) bool {
				distances[r][c] -= rowMins[r]
				return true
			})

			return 2
		}

		step2 := func() int {
			/*
				Step 2:  Find a zero (Z) in the resulting matrix.
				If there is no starred zero in its row or column, star Z.
				Repeat for each element in the matrix.
				Go to Step 3.
			*/
			cols, rows := make([]bool, L), make([]bool, L)
			visitAll(distances, func(r, c, d int) bool {
				if 0 == d {
					if !rows[r] && !cols[c] {
						starredZeros = append(starredZeros, cell{r: r, c: c})
					}
					rows[r] = true
					cols[c] = true
				}

				return true
			})

			return 3
		}

		step3 := func() int {
			/*
				Step 3:  Cover each column containing a starred zero.
				 If K columns are covered,
				  the starred zeros describe a complete set of unique assignments.
				In this case, Go to DONE, otherwise, Go to Step 4.
			*/

			c := 0
			for _, z := range starredZeros {
				if !coveredCols[z.c] {
					coveredCols[z.c] = true
					c++
				}
			}
			if c >= L {
				return 0
			}

			return 4
		}

		step4 := func() int {
			/*
				Step 4: Find a noncovered zero and prime it.
				If there is no starred zero in the row containing this primed zero,
				 Go to Step 5.
				 Otherwise, cover this row and uncover the column containing the starred zero. Continue in this manner until there are no uncovered zeros left.
				Save the smallest uncovered value and Go to Step 6.
			*/
			gotoStep5 := false
			starredRows := make([]int, L)
			for i := range starredRows {
				starredRows[i] = -1
			}
			for _, z := range starredZeros {
				starredRows[z.r] = z.c
			}

			visitAll(distances, func(r, c, d int) bool {
				if 0 == d && !coveredCols[c] && !coveredRows[r] {
					primedZeros = append(primedZeros, cell{r: r, c: c})
					starredCol := starredRows[r]
					if starredCol > -1 {
						coveredRows[r] = true
						coveredCols[starredCol] = false
					} else {
						gotoStep5 = true
						return false
					}
				}
				return true
			})
			if gotoStep5 {
				return 5
			}

			return 6
		}

		step5 := func() int {
			/*
			 Step 5: Construct a series of alternating primed and starred zeros as follows.
			 Let Z0 represent the uncovered primed zero found in Step 4.
			 Let Z1 denote the starred zero in the column of Z0 (if any).
			 Let Z2 denote the primed zero in the row of Z1 (there will always be one).
			 Continue until the series terminates at a primed zero that has no starred zero in its column.
			 Unstar each starred zero of the series, star each primed zero of the series,
			  erase all primes and uncover every line in the matrix.
			 Return to Step 3.
			*/
			/*
				fmt.Println(starredZeros)
				fmt.Println(primedZeros)
			*/
			zeroesToUnstar := map[int]bool{}
			zeroesToStar := map[int]bool{}
			for i0, z0 := range primedZeros {
				found := false
				var i1, i2 int
				var z1, z2 cell
				for i1, z1 = range starredZeros {
					zeroesToStar[i0] = true
					found = z0.c == z1.c
					if found {
						zeroesToUnstar[i1] = true
						break
					}
				}

				if !found {
					break
				}

				for i2, z2 = range primedZeros {
					if z2.r == z1.r {
						zeroesToStar[i2] = true
						break
					}
				}

			}

			newStarred := []cell{}
			for i, z := range starredZeros {
				if !zeroesToUnstar[i] {
					newStarred = append(newStarred, z)
				}
			}

			for i, z := range primedZeros {
				if zeroesToStar[i] {
					newStarred = append(newStarred, z)
				}
			}

			starredZeros = make([]cell, len(newStarred))
			copy(starredZeros, newStarred)
			primedZeros = []cell{}
			/*
				fmt.Println(starredZeros)
			*/
			for i := range coveredCols {
				coveredCols[i] = false
				coveredRows[i] = false
			}

			return 3
		}

		step6 := func() int {
			/*
				Step 6: Add the value found in Step 4 to every element of each covered row,
					and subtract it from every element of each uncovered column.
				Return to Step 4 without altering any stars, primes, or covered lines
			*/

			smallestUncovered := max
			visitAll(distances, func(r, c, d int) bool {
				if !coveredCols[c] && !coveredRows[r] && d < smallestUncovered {
					smallestUncovered = d
				}

				return true
			})

			visitAll(distances, func(r, c, d int) bool {
				if coveredRows[r] {
					if distances[r][c] < max {
						distances[r][c] += smallestUncovered
					}
				}

				if !coveredCols[c] {
					distances[r][c] -= smallestUncovered
				}

				return true
			})
			return 4
		}

		steps := [7]munkersStep{nil, step1, step2, step3, step4, step5, step6}

		step := 1
		for step > 0 {
			//	fmt.Println("step ", step)
			step = steps[step]()
			printMatrix()
		}

		return starredZeros
	}

	prepare(distances)
	zeroes := munkers()

	res := -1

	for _, z := range zeroes {
		if orig[z.r][z.c] > res {
			res = orig[z.r][z.c]
		}
	}

	return res
}
