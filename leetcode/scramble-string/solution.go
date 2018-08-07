package main

import (
	"fmt"
	"strings"
)

func isScramble(s1 string, s2 string) bool {
	sum := 0
	for i := 0; i < len(s1); i++ {
		sum += i - strings.Index(s2, string(s1[i]))
	}

	return 0 == sum
}

func main() {
	fmt.Println(isScramble("great", "rgeat"))
	fmt.Println(isScramble("abcde", "caebd"))
}
