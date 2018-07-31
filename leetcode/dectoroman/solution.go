package main

import (
	"fmt"
	"sort"
)

func intToRoman(num int) string {
	result := ""

	digits := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	keys := []int{}
	for k := range digits {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for i := len(keys) - 1; i >= 0; i-- {
		key := keys[i]
		for ; num >= key; num -= key {
			result += digits[key]
		}
	}
	/*
		for i := len(keys) - 2; i >= 0; i-- {
			digit := digits[keys[i]]
			nextDigit := digits[keys[i+1]]

			result = strings.Replace(result, strings.Repeat(digit, 4), digit+nextDigit, -1)
		}
	*/
	return result
}

func main() {
	fmt.Println(intToRoman(19), intToRoman(1994), intToRoman(3999))
}
