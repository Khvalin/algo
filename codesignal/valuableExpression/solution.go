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
	const M = 1000000007

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

	extractTop2ForProduct := func(digits *[]int, total int) (int, int) {
		res := [2]int{}

		for i := 0; i < total; i++ {
			index := 0
			if res[1] < res[0] {
				index = 1
			}

			res[index] *= 10
			res[index] += (*digits)[i]

			res[index] %= M
		}
		*digits = (*digits)[total:]

		return res[1], res[0]
	}

	var result int

	digits, ops := createQueques(hand)

	if ops[0] == multiply {
		topCount := len(digits) - len(ops) + 1
		ops = ops[1:]

		if digits[0] > 0 && digits[1] > 0 {
			//optimistic product
			//len1, len2 := int(math.Ceil(float64(topCount)/2)), int(math.Floor(float64(topCount)/2))
			a, b := extractTop2ForProduct(&digits, topCount)
			result = (a * b) % M

			fmt.Println(digits)
			fmt.Println(a, b, result)
		} else {
			// pessimistic product
			//a, b := digits[len(digits)-1], digits[len(digits)-2]
			result, digits = digits[0], append(digits[1:len(digits)-2], 0)
		}

	} else {
		topCount := len(digits) - len(ops)
		for i := 0; i < topCount; i++ {
			result *= 10
			result += digits[i]

			result %= M
		}
		digits = digits[topCount:]
	}

	for i := 0; i < len(ops); i++ {
		koef := 1
		if ops[i] == minus {
			koef = -1
		}

		result = (result + koef*digits[i]) % M
	}

	return result
}

func main() {
	fmt.Println(valuableExpression([]string{"1", "0", "0", "*", "-"}))
	fmt.Println(valuableExpression([]string{"9", "7", "9", "0", "+", "+", "8", "-", "+", "-", "+", "+", "9", "7", "3", "7", "*", "+", "6", "7", "3", "+", "-", "5", "2", "5", "5", "6", "0", "6", "+", "+", "-", "8", "8", "-", "8", "2", "1", "2", "1", "1"}))

	fmt.Println(valuableExpression([]string{"4", "9", "+", "8", "1", "5", "-", "+", "5"}))

	fmt.Println(valuableExpression([]string{"8", "4", "2", "7", "5", "*"}))
	fmt.Println(valuableExpression([]string{"1", "5", "3", "*", "1", "+"}))

	fmt.Println(valuableExpression([]string{"9", "1", "3", "4", "*", "-", "+"}))

	fmt.Println(valuableExpression([]string{"9", "1", "9", "1", "*"}))

}
