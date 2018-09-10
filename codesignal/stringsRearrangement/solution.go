package main

import (
	"fmt"
)

func stringsRearrangement(inputArray []string) bool {
	//TODO: add memorization maybe?
	areArranged := func(s1, s2 string) bool {
		diff := 0
		for i := range s1 {
			if s1[i] != s2[i] {
				diff++
			}
		}

		return 1 == diff
	}

	var solve func(start, end string, input []string) bool

	// TODO: use indices array instead of strings?
	solve = func(start, end string, input []string) bool {
		result := len(input) == 0
		for i, s := range input {
			if areArranged(end, s) {
				a := []string{}
				a = append(a, input[:i]...)
				a = append(a, input[i+1:]...)
				result = result || solve(start, s, a)
			}
			if start != end && areArranged(start, s) {
				a := []string{}
				a = append(a, input[:i]...)
				a = append(a, input[i+1:]...)
				result = result || solve(s, end, a)
			}
		}

		return result
	}

	return solve(inputArray[0], inputArray[0], inputArray[1:])
}

func main() {
	fmt.Println(true, stringsRearrangement([]string{"ab", "bb", "aa"}))
	fmt.Println(false, stringsRearrangement([]string{"aba", "bbb", "bab"}))
	fmt.Println(true, stringsRearrangement([]string{"abc", "bef", "bcc", "bec", "bbc", "bdc"}))
}
