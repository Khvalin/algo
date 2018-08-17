package main

import (
	"fmt"
	"strings"
)

func findSubstring(s string, words []string) []int {
	//LEN := (len(words[0]))
	result := []int{}

	wordsMap := map[string]bool{}

	for index := strings.Index(s, words[0]); index > -1; index = strings.Index(s, words[0]) {
		for i := 1; i < len(words); i++ {
			wordsMap[words[i]] = true
		}

	}
	return result
}

func main() {
	fmt.Println("barfoothefoobarman", []string{"foo", "bar"})
}
