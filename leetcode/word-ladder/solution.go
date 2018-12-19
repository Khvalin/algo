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
		last[i] = L
		connected[i] = make([]bool, L)
	}

	for i := 0; i < L; i++ {
		for j := i; j < L; j++ {
			connected[i][j] = 1 == diff(wordList[i], wordList[j])
			connected[j][i] = connected[i][j]
			if last[i] > j {
				last[i] = j
			}
		}
	}
	if connected[0][1] {
		return 2
	}
	//	fmt.Println(connected[0])

	res := len(wordList) + 10
	origLast := make([]int, L)
	copy(origLast, last)

	stack, used := []int{1}, make(map[int]bool)
	for len(stack) > 0 {
		start := stack[len(stack)-1]

		//	next := L + 1

		next := last[start] + 1
		for next < L && (!connected[start][next] || used[next]) {
			next++
		}
		/*
			ans := ""
			for _, v := range stack {
				ans += wordList[v] + " "
			}
			fmt.Println(ans) */

		last[start] = next

		if next >= L || len(stack) > res {
			//rewind
			l := len(stack) - 1
			for l >= 0 && last[stack[l]] >= L { // TODO : l >= 0
				//	last[stack[l]] = origLast[stack[l]]
				used[stack[l]] = false
				l--
			}
			stack = stack[:l+1]
			// stack = stack[:l+1]
		} else {
			chainFound := connected[next][0]
			if chainFound {
				if res > len(stack) {
					res = len(stack)
				}
			} else {
				used[next] = true
				stack = append(stack, next)
			}
		}
	}

	if res > len(wordList) {
		return 0
	}

	return res + 2
}
