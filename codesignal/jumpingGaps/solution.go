package main

import "fmt"

import "strings"

type cell struct {
	r, c   int
	dist   int
	isWall bool
}

//Board : cell matrix
type Board [][]cell

const max = 1 << 31

func (c cell) String() string {
	//	return fmt.Sprintf("{r: %v, c: %v}", c.r, c.c)
	d := c.dist
	res := fmt.Sprintf("%v", d)
	if d == max {
		res = "*"
	}
	if c.isWall {
		res = "W"
	}
	return res
}

func prepare(stage *[]string) (cell, cell, Board) {
	st := *stage
	L := len(st[0])
	empty, wall := strings.Repeat(" ", L), strings.Repeat("#", L)
	st = append([]string{wall, empty}, st...)
	st = append(st, wall)

	b := make(Board, len(st))
	var end, start cell

	for i, s := range st {
		b[i] = make([]cell, L)
		for j, ch := range s {
			d := max
			if 'S' == ch {
				d = 0
			}
			b[i][j] = cell{r: i, c: j, dist: d, isWall: ('#' == ch)}
			if 'E' == ch {
				end = b[i][j]
			}
			if 'S' == ch {
				start = b[i][j]
			}
		}
	}

	return start, end, b
}

func dfs(start, end cell, board Board) int {
	q := []cell{start}
	//	N := len(board)
	M := len(board[0])
	for len(q) > 0 {
		s := q[0]
		q = q[1:]
		//fmt.Println(s.r, s.c)
		for i := 0; i <= 3; i++ {
			x := s.r - i

			if x < 0 || board[x][s.c].isWall {
				break
			}
			for k := -1; k <= 1; k += 2 {
				for j := 1; j <= 3; j++ {
					y := s.c + j*k
					//fmt.Println("  ", x, y)
					if y < 0 || y >= M {
						break
					}

					if board[x][y].isWall {
						break
					}

					r := x
					//falling down:
					for !board[r+1][y].isWall {
						r++
					}

					d := s.dist + 1
					if board[r][y].dist > d && r > 0 { // 0 === top frame border
						board[r][y].dist = d
						q = append(q, board[r][y])
					}

				}
			}
		}
	}
	/*
		for _, a := range board {
			fmt.Println(a)
		}
	*/
	res := board[end.r][end.c].dist
	if max == res {
		res = -1
	}

	return res
}

func jumpingGaps(stage []string) int {
	start, end, board := prepare(&stage)

	return dfs(start, end, board)
}

func main() {
	//
}
