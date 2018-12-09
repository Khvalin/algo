package wordLadder

import (
	"fmt"
)

type Item struct {
	used map[int]bool
	path []int
}

func (item *Item) String() string {
	return fmt.Sprintf("%b", item.path)
}

func copyItem(i *Item) *Item {
	res := &Item{path: []int{}, used: map[int]bool{}}
	if i != nil {
		for key, value := range i.used {
			res.used[key] = value
		}
		res.path = append(i.path, res.path...)
	}
	return res
}

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

	q := []Item{Item{path: []int{-1}, used: map[int]bool{}}}

	for len(q) > 0 {
		v := q[0]
		q = q[1:]

		tmp := ""
		for _, n := range v.path {
			if n > -1 {
				tmp += wordList[n] + " "
			}
		}

		fmt.Println(tmp)

		last := v.path[len(v.path)-1]
		if last > -1 && wordList[last] == endWord {
			res = len(v.used)
			break
		}

		for _, n := range T[last] {
			if !v.used[n] {
				newItem := *copyItem(&v)
				newItem.path = append([]int{n}, v.path...)
				newItem.used[n] = true

				q = append([]Item{newItem}, q...)
			}
		}
	}

	if res > len(wordList) {
		return 0
	}
	return res
}
