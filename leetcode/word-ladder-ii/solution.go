package wordLadder

import (
//"fmt"
)

func ladderLength(beginWord string, endWord string, wordList []string) (int, [][]int) {
	endWordIndex := -1
	for i, w := range wordList {
		if w == endWord {
			endWordIndex = i
		}
	}

	if endWordIndex == -1 {
		return 0, nil
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

	//sort.Strings(wordList)
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

	connected, pathLen := make([][]int, L), make([]int, L)
	for i := 0; i < L; i++ {
		pathLen[i] = L + 1
		connected[i] = []int{}
	}

	for i := 0; i < L; i++ {
		for j := i + 1; j < L; j++ {
			if 1 == diff(wordList[i], wordList[j]) {
				connected[j] = append(connected[j], i)
				connected[i] = append(connected[i], j)
			}
		}
	}

	pathLen[1] = 1

	q := []int{1}
	neighbours := make([]int, L)
	for len(q) > 0 {
		start := q[0]
		q = q[1:]

		j := 0
		for _, v := range connected[start] {
			if pathLen[v] > pathLen[start]+1 {
				pathLen[v] = pathLen[start] + 1
				neighbours[j] = v
				j++
			}
		}
		q = append(q, neighbours[:j]...)

	}

	res := pathLen[0]
	//fmt.Printf("")

	if res > len(wordList) {
		return 0, nil
	}

	return res, connected
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	L, res := len(wordList), [][]string{}

	minLen, connected := ladderLength(beginWord, endWord, wordList)
	visited := make([]bool, L)
	last := make([]int, L)
	for i := range last {
		last[i] = 1
	}

	stack := []int{1}
	for len(stack) > 0 {
		start := stack[len(stack)-1]

		i := last[start] + 1
		v := connected[start]
		for i < len(v) && !visited[v[i]] {
			i++
		}

		if i < len(v) {
			if len(stack) == minLen {
				//
			}
		} else {
			// rewind
		}
	}
	return res
}
