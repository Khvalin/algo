package main

import "fmt"

func letterCombinations(digits string) []string {
	LENGTH := len(digits)

	if LENGTH == 0 {
		return []string{}
	}

	var digitsMap = map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	toLetters := func(mapping []byte) string {
		result := ""

		for i := 0; i < LENGTH; i++ {
			result += string(digitsMap[digits[i]][mapping[i]])
		}

		return result
	}

	var result []string

	var variants = make([]byte, LENGTH+1)

	for i := 0; i < LENGTH; i++ {
		variants[i] = 0
	}

	for variants[LENGTH] == 0 {
		result = append(result, toLetters(variants))
		cur := 0

		overflow := func() bool {
			return int(variants[cur]) >= len(digitsMap[digits[cur]])
		}

		variants[cur]++

		for cur < LENGTH && overflow() {
			variants[cur] = 0
			cur++
			variants[cur]++
		}

	}

	return result
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
}
