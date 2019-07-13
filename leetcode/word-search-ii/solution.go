package main

import (
	"fmt"
)

func exist(board [][]byte, word string) bool {
	var visit func(x, y int, word string) bool
	visit = func(x, y int, word string) bool {
		ch, w := word[0], word[1:]

		if board[x][y] != ch {
			return false
		}

		if len(w) == 0 {
			return true
		}

		board[x][y] = ' '
		r := false
		if x > 0 {
			r = visit(x-1, y, w)
		}

		if !r && y > 0 {
			r = visit(x, y-1, w)
		}

		if !r && x < len(board)-1 {
			r = visit(x+1, y, w)
		}

		if !r && y < len(board[x])-1 {
			r = visit(x, y+1, w)
		}

		board[x][y] = ch

		return r
	}

	for i := range board {
		for j := range board[i] {
			if visit(i, j, word) {
				return true
			}

		}
	}

	return false
}

func findWords(board [][]byte, words []string) []string {
	var res []string

	for _, w := range words {
		if exist(board, w) {
			res = append(res, w)
		}
	}

	return res
}

func main() {
	BOARD := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'A', 'B', 'C', 'S'},
		[]byte{'C', 'B', 'E', 'E'},
		[]byte{'A', 'D', 'E', 'E'},
	}

	fmt.Println(findWords(BOARD, []string{"ABBA", "EEEEE"}))
}
