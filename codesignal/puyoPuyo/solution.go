package main

import (
	"fmt"
	"strings"
)

const NROWS = 15
const NCOLS = 6

type TBoard [NROWS][NCOLS]rune

func puyoPuyo(field []string, nextPiece string) int {
	board := new(TBoard)

	/*
		createChainMap := func () [][]int{} {

		}
	*/
	for i, j := len(field)-1, 0; i >= 0; i, j = i-1, j+1 {
		runes := []rune(strings.Replace(field[i], " ", "\000", -1))
		copy(board[j][:], runes)
	}

	getScore := func(board TBoard, cols []int) int {
		for _, col := range cols {
			row := 0
			for ; board[row][col] > 0; row++ {
			}

			for i := row; i < NROWS; i++ {
				if board[i][col] > 0 {
					board[row][col] = board[i][col]
					board[i][col] = 0
					row++
				}
			}
		}

		counts := [NROWS][NCOLS]int{}
		for _, col := range cols {
			var row int
			for row = 0; board[row][col] > 0; row++ {
				ch := board[row][col]
				if counts[row][col] == 0 {
					//
				}
			}

		}

		return 0
	}

	pieces := []string{nextPiece, string(nextPiece[1]) + string(nextPiece[0])}

	for _, piece := range pieces {
		for i := 0; i < NCOLS; i++ {
			board[NROWS-1][i], board[NROWS-2][i] = rune(piece[0]), rune(piece[1])
			getScore(*board, []int{i})

			board[NROWS-1][i], board[NROWS-2][i] = 0, 0
		}
	}

	return 0
}

func main() {

	input1 := []string{
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"  Y   ",
		"  Y   ",
		"  Y   ",
		"  B   ",
		"  BR  ",
		"GYYBR ",
		"GGYBR ",
	}

	fmt.Println(puyoPuyo(input1, "YG"))
}
