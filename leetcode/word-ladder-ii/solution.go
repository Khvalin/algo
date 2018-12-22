package wordLadder

import (
//"fmt"
)

func ladderLength(beginWord string, endWord string, wordList []string) (int, [][]int, []string) {
	endWordIndex := -1

	for i, w := range wordList {
		if w == endWord {
			endWordIndex = i
		}
	}

	if endWordIndex == -1 {
		return 0, nil, wordList
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
		return 0, nil, wordList
	}

	return res, connected, wordList
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	minLen, connected, wordList := ladderLength(beginWord, endWord, wordList)
	//fmt.Println(wordList)
	L, res := len(wordList), [][]string{}

	visited := make([]bool, L)
	last := make([]int, L)
	for i := range last {
		last[i] = -1
	}
	last[0] = L + 1
	visited[1] = true

	stack := []int{1}
	for len(stack) > 0 {
		start := stack[len(stack)-1]

		i := last[start] + 1
		v := connected[start]

		for i < len(v) && v[i] < len(visited) && visited[v[i]] {
			i++
		}
		last[start] = i

		if i < len(v) && len(stack) <= minLen {
			next := v[i]

			//	fmt.Println(next)
			visited[next] = true
			stack = append(stack, next)

			if len(stack) == minLen && next == 0 {
				path := []string{}
				for _, k := range stack {
					path = append(path, wordList[k])
				}
				res = append(res, path)
			}

		} else {
			// rewind
			l := len(stack) - 1
			for l >= 0 && last[stack[l]] > len(connected[stack[l]])-1 {
				visited[stack[l]] = false
				last[stack[l]] = -1
				l--
			}
			stack = stack[:l+1]
		}
	}
	//	fmt.Println(res)

	return res
}
