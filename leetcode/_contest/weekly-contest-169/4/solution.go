package main

import "fmt"

func isSolvable(words []string, result string) bool {
	dict := map[rune]int{}
	abc := []rune{}

	voc := words
	voc = append(voc, result)

	c := 0
	for _, w := range voc {
		for _, ch := range w {
			if _, found := dict[ch]; found {
				continue
			}
			abc = append(abc, rune(ch))

			dict[ch] = c
			c++
		}
	}

	trySolve := func(m map[rune]int) bool {
		sum := 0
		for i, w := range voc {
			if len(w) > 0 && m[rune(w[0])] == 0 {
				return false
			}
			n := 0
			for _, ch := range w {
				n *= 10
				n += m[rune(ch)]
			}

			if i == len(voc)-1 {

				return n == sum
			}
			sum += n
		}

		return true
	}

	for dict[abc[len(abc)-1]] < 10 {
		dict[abc[0]]++

		for i := 0; i < len(abc)-2; i++ {
			dict[abc[i+1]] += dict[abc[i]] / 10
			dict[abc[i]] %= 10
		}

		if dict[abc[len(abc)-1]] < 10 && trySolve(dict) {
			return true
		}

	}

	return false
}

func main() {
	//fmt.Println(isSolvable([]string{"AA", "BB"}, "DEC"))
	//fmt.Println(isSolvable([]string{"SEND", "MORE"}, "MONEY"))
	fmt.Println(isSolvable([]string{"THIS", "IS", "TOO"}, "FUNNY"))
	fmt.Println(isSolvable([]string{"LEET", "CODE"}, "POINT"))
}
