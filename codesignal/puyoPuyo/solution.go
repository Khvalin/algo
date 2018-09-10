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

	getScore := func(board TBoard) int {

		board[NROWS-1][2] = 'z'

		return 0
	}

	go getScore(*board)
	fmt.Println(board)
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
