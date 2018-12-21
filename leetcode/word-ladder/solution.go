package wordLadder

import (
//	"fmt"
)

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
			if (1 == diff(wordList[i], wordList[j])) {
				connected[j] = append(connected[j], i)
				connected[i] = append(connected[i], j)
			}
		}
	}

	//	fmt.Println(connected[0])

	pathLen[1]= 1

	q := []int{1}
	neighbours := make([]int, L)
	for len(q) > 0 {
		start := q[0]
		q = q[1:]

		//	next := L + 1

		j := 0
		for _, v := range connected[start] {
			if (pathLen[v] > pathLen[start]+1) {
				pathLen[v] = pathLen[start]+1
				neighbours[j] = v
				j++
			}
		}
		q = append(q, neighbours[:j]...)

	}
	

	res := pathLen[0]
	//fmt.Println(pathLen)

	if res > len(wordList) {
		return 0
	}

	return res 
}
