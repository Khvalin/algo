package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

func findMinStep(board string, hand string) int {
	const THREE = 3
	type Zuma struct {
		board string
		hand  string
	}
	var solve func(variant Zuma, step int) int
	var reduce func(board string) string

	reduce = func(board string) string {
		fmt.Println(" reducing : ", board)
		count := 0
		prev := rune(0)

		for i := 0; i < len(board); i++ {
			if prev != rune(board[i]) {
				count = 0
			}
			count++

			if count == THREE {
				re := regexp.MustCompile(string(prev) + "{3,}")
				return reduce(re.ReplaceAllString(board, ""))

			}

			prev = rune(board[i])
		}

		return (board)
	}

	solve = func(variant Zuma, step int) int {
		fmt.Println(" ", step, " ", variant)

		result := math.MaxInt32

		genVariants := func(variant Zuma) []Zuma {
			handMap := map[byte]int{}
			for i := 0; i < len(variant.hand); i++ {
				_, found := handMap[variant.hand[i]]
				if !found {
					handMap[variant.hand[i]] = 0
				}

				handMap[variant.hand[i]]++
			}

			count := 0

			board := variant.board
			prev := rune(0)
			variants := []Zuma{}
			for i := 0; i < len(board); i++ {
				if prev != rune(board[i]) {
					count = 0
				}
				count++

				ocs, found := handMap[board[i]]
				if found && count+ocs >= THREE {
					insertCount := THREE - count
					newBoard := reduce(board[:i] + strings.Repeat(string(board[i]), insertCount) + board[i:])
					newHand := strings.Replace(variant.hand, string(board[i]), "", insertCount)

					variants = append(variants, Zuma{board: newBoard, hand: newHand})
				}

				prev = rune(board[i])
			}

			return variants
		}

		if len(variant.board) == 0 {
			return step + 1
		}

		variants := genVariants(variant)

		if len(variants) > 0 {
			for i := 0; i < len(variants); i++ {
				minSteps := solve(variants[i], step+1)

				if minSteps < result {
					result = minSteps
				}
			}
		}

		return result
	}

	result := solve(Zuma{board, hand}, 0)

	if result < math.MaxInt32 {
		return result
	}

	return -1
}

func main() {
	//fmt.Println(findMinStep("WRRBBW", "WR"))
	fmt.Println(findMinStep("WWRRBBWW", "WRBRW"))
	//fmt.Println(findMinStep("G", "GGGGG"))
	//fmt.Println(findMinStep("RBYYBBRRB", "YRBGB"))
}
