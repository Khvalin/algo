package main

import (
	"fmt"
	"strings"
)

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	num, lastint, total := 0, 0, 0
	for i := 0; i < len(s); i++ {
		char := s[len(s)-(i+1) : len(s)-i]
		num = roman[char]
		if num < lastint {
			total = total - num
		} else {
			total = total + num
		}
		lastint = num
	}
	return total
}

func solve(a, b int) int {
	if a < b {
		return -1
	}

	if a > b {
		return 1
	}

	return 0
}

func main() {
	num1, num2 := "", ""
	fmt.Scanln(&num1, &num2)

	arr1 := strings.Split(num1, "|")
	arr2 := strings.Split(num2, "|")

	if len(arr1) != len(arr2) {
		fmt.Println(solve(len(arr1), len(arr2)))
		return
	}

	for i := 0; i < len(arr1); i++ {
		a := romanToInt(arr1[i])
		b := romanToInt(arr2[i])

		res := solve(a, b)
		if res != 0 {
			fmt.Println(res)
			return
		}
	}

	fmt.Println(0)
}
