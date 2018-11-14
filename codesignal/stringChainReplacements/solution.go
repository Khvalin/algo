package main

import (
	"fmt"
	"strings"
)

func stringChainReplacements(stringArray []string) int {
	chain := make([]string, len(stringArray))
	for i, s := range stringArray {
		chain[i] = string(s[0])
		if i == 0 {
			chain[i] = string(s[len(s)-1])
			continue
		}

		if i == len(stringArray)-1 {
			chain[i] = string(s[0])
			continue
		}

		if len(s) > 1 {
			chain[i] += string(s[len(s)-1])
		}
	}

	M := map[string]int{}

	var count func(start string, subChain []string, index int) int
	count = func(start string, subChain []string, index int) int {
		if len(subChain) < 1 {
			return 0
		}
		hash := start + " " + strings.Join(subChain, " ")
		c, ok := M[hash]
		if !ok {
			second := subChain[0]

			c = count(second, subChain[1:], index+1)
			//TODO: backwards changes count

			if start[len(start)-1] != second[0] {
				second = string(start[len(start)-1])
				alt := count(second, subChain[1:], index+1) + 1
				if alt < c {
					c = alt
				}
			}
			M[hash] = c
		}

		fmt.Println(hash, c)
		return c
	}

	res := count(chain[0], chain[1:], 0)
	//fmt.Println(M)
	return res
}

func main() {
	/*
		fmt.Println(stringChainReplacements([]string{"abc",
			"def",
			"ghi"}))

		fmt.Println(stringChainReplacements([]string{"q",
			"p",
			"r",
			"qrpst",
			"tv",
			"a"}), 3)

		fmt.Println(stringChainReplacements([]string{"a", "bb", "a"}), 2)
	*/
	fmt.Println(stringChainReplacements([]string{"xt",
		"x",
		"fqrhsqxwt",
		"tontq",
		"nk",
		"knplki",
		"i",
		"x",
		"g",
		"gifguc",
		"e"}), 7)
}
