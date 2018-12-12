package main

import "fmt"
import "sort"

func longestWordLadder(startWord string, goalWord string, usableWords []string) []string {
	sort.Strings(usableWords)
	usableWords = append([]string{goalWord, startWord}, usableWords...)

	L := len(usableWords)

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
			connected[i][j] = 1 == diff(usableWords[i], usableWords[j])
			connected[j][i] = connected[i][j]
		}
	}

	//fmt.Println(connected)

	var resIndices []int
	stack, used := []int{1}, make(map[int]bool)
	for len(stack) > 0 {
		index := stack[len(stack)-1]
		if connected[index][0] && len(stack) > len(resIndices) {
			resIndices = make([]int, len(stack))
			copy(resIndices, stack)
			resIndices = append(resIndices, 0)
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

	res := make([]string, len(resIndices))
	for i, index := range resIndices {
		res[i] = usableWords[index]
	}
	return res
}

func main() {
	fmt.Println(longestWordLadder("code", "dope", []string{"cone", "bonk", "none", "dole", "tune", "node", "mode", "nope", "mole"}))
	beginWord := "qa"
	endWord := "sq"
	wordList := []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"}
	res := longestWordLadder(beginWord, endWord, wordList)

	fmt.Println(res)
}
