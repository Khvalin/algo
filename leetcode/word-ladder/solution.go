package wordLadder

import "sort"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	endWordIndex := -1
	for i, w := range wordList {
		if w == endWord {
			endWordIndex = i
		}
	}

	if endWordIndex == -1 {
		return 0
	}

	wordList[endWordIndex] = wordList[len(wordList)-1]
	wordList = wordList[:len(wordList)-1]

	beginWordIndex := -1
	for i, w := range wordList {
		if w == beginWord {
			beginWordIndex = i
			break
		}
	}

	if beginWordIndex > -1 {
		wordList[beginWordIndex] = wordList[len(wordList)-1]
		wordList = wordList[:len(wordList)-1]
	}

	sort.Strings(wordList)
	wordList = append([]string{endWord, beginWord}, wordList...)

	L := len(wordList)

	diff := func(a, b string) (c byte) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				c++
				if c > 1 {
					return
				}
			}
		}
		return
	}

	connected, last := make([][]bool, L), make([]int, L)
	for i := 0; i < L; i++ {
		last[i] = 1
		connected[i] = make([]bool, L)
	}

	for i := 0; i < L; i++ {
		for j := i; j < L; j++ {
			connected[i][j] = 1 == diff(wordList[i], wordList[j])
			connected[j][i] = connected[i][j]
		}
	}

	res := len(wordList) + 10

	stack, used := []int{1}, make(map[int]bool)
	for len(stack) > 0 {
		index := stack[len(stack)-1]
		if connected[index][0] && res > len(stack) {
			res = len(stack)
		}

		i := last[index] + 1
		for i < L && (!connected[i][index] || used[i]) {
			i++
		}
		last[index] = i

		if i < L {
			used[index] = true
			stack = append(stack, i)
		} else {
			l := len(stack) - 1
			for l > 0 && last[stack[l]] >= L {
				last[stack[l]] = 1
				used[stack[l]] = false
				l--
			}
			stack = stack[:l]
		}
	}

	if res > len(wordList) {
		return 0
	}

	return res + 1
}
