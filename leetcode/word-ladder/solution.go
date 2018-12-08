package wordLadder

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	canTransform := func(str1, str2 string) bool {
		r := 0
		for i := 0; i < len(str1); i++ {
			if str1[i] != str2[i] {
				r++
			}
		}

		return r == 1
	}

	T := map[int][]int{}
	for i, s := range wordList {
		if canTransform(beginWord, s) {
			T[-1] = append(T[-1], i)
		}
	}

	for i := 0; i < len(wordList); i++ {
		for j := i + 1; j < len(wordList); j++ {
			if canTransform(wordList[i], wordList[j]) {
				T[j] = append(T[j], i)
				if wordList[i] != endWord {
					T[i] = append(T[i], j)
				}
			}
		}
	}

	res := len(wordList) + 10

	stack := []int{-1}
	used := map[int]bool{}
	last := map[int]int{}

	for len(stack) > 0 {
		c := stack[len(stack)-1]

		i := last[c]

		for i < len(T[c]) && used[T[c][i]] {
			i++
		}
		if i < len(T[c]) {
			last[c] = i + 1
			n := T[c][i]
			stack = append(stack, n)
			used[n] = true
		} else {
			if c >= 0 && wordList[c] == endWord && len(stack) < res {
				res = len(stack)
			}
			used[c] = false
			last[c] = 0

			stack = stack[:len(stack)-1]
		}
		fmt.Println(stack)
	}

	if res > len(wordList) {
		return 0
	}
	return res
}
