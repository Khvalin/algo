package main

import (
	"fmt"
	"sort"
	"strconv"
)

const multiply = 10
const plus = 20
const minus = 30

func valuableExpression(hand []string) int {
	createQueques := func(hand []string) ([]int, []int) {
		digits, ops := []int{}, []int{}
		for i := 0; i < len(hand); i++ {
			if hand[i] >= "0" && hand[i] <= "9" {
				d, _ := strconv.Atoi(hand[i])
				digits = append(digits, d)
			}
			if hand[i] == "*" {
				ops = append(ops, multiply)
			}
			if hand[i] == "+" {
				ops = append(ops, plus)
			}
			if hand[i] == "-" {
				ops = append(ops, minus)
			}
		}

		sort.Sort(sort.Reverse(sort.IntSlice(digits)))
		sort.Ints(ops)

		return digits, ops
	}

	var result int
	digits, ops := createQueques(hand)
	fmt.Println(digits, ops)

	return result
}

func main() {
	fmt.Println(valuableExpression([]string{"9",
		"1",
		"3",
		"4",
		"*",
		"-",
		"+"}))
}
