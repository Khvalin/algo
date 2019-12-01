package main

import "fmt"

func tictactoe(moves [][]int) string {
	B := [3][3]rune{}
	res := "Pending"
	aToMove := true

	for j, m := range moves {
		ch := 'A'
		if !aToMove {
			ch = 'B'
		}

		row, col := m[0], m[1]
		B[row][col] = ch

		for i := 0; i < 3; i++ {
			if (B[0][i] == B[1][i] && B[1][i] == B[2][i] && B[0][i] == ch) ||
				(B[i][0] == B[i][1] && B[i][0] == B[i][2] && B[i][0] == ch) ||

				(B[0][0] == B[1][1] && B[0][0] == B[2][2] && B[0][0] == ch) ||
				B[0][2] == B[1][1] && B[0][2] == B[2][0] && B[2][0] == ch {

				return string(ch)
			}
		}

		if j == 8 {
			res = "Draw"
		}

		aToMove = !aToMove
	}

	return res
}

func main() {
	moves := [][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}
	r := tictactoe(moves)
	fmt.Println(r)

	moves = [][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}}
	r = tictactoe(moves)
	fmt.Println(r)
}
