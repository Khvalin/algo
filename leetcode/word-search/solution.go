package main

import (
	"fmt"
)

func exist(board [][]byte, word string) bool {
	var ch byte
	ch, word = word[0], word[1:]
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if ch == board[i][j] {
				if len(word) == 0 {
					// TODO
					return true
				}
				// TODO
				board[i][j] = ' '
				if exist(board, word) {
					return true
				}
				board[i][j] = ch

			}
		}
	}

	return false
}

func main() {
	BOARD := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'A', 'B', 'C', 'S'},
		[]byte{'C', 'B', 'E', 'E'},
		[]byte{'A', 'D', 'E', 'E'},
	}

	fmt.Println(exist(BOARD, "ABBAA"))
}
