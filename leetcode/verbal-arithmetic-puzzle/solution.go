package main

import (
	"fmt"
)

func isSolvable(words []string, result string) bool {
	decypher := func(dict [95]byte, rev [10]rune) bool {
		sum := 0
		for _, w := range words {
			n := 0
			for _, ch := range w {
				n *= 10
				n += int(dict[ch])
			}
			sum += n
		}

		for i := len(result) - 1; i >= 0; i-- {
			d, ch := byte(sum%10), rune(result[i])
			sum /= 10
			if dict[ch] == 250 {
				if rev[d] > 0 {
					return false
				}
				rev[d] = ch
				dict[ch] = d
			}
			if dict[ch] != d {
				return false
			}
		}

		if len(result) > 1 && dict[rune(result[0])] == 0 {
			return false
		}

		return true
	}

	l := 0
	abc := [95]bool{}
	noZeros := [95]bool{}
	if len(result) > 1 {
		noZeros[result[0]] = true
	}

	for _, w := range words {
		for i, ch := range w {
			if len(w) > 1 && i == 0 {
				noZeros[ch] = true
			}
			if !abc[ch] {
				l++
			}
			abc[ch] = true
		}
	}

	dict := [95]byte{}
	for i := range dict {
		dict[i] = 250
	}
	res := false

	var next func(min byte, rev [10]rune)
	next = func(min byte, rev [10]rune) {
		if res {
			return
		}
		count := 0
		for j := 0; j < len(rev); j++ {
			if rev[j] == 0 {
				for i := min; i <= 'Z'; i++ {
					if abc[i] && (j > 0 || !noZeros[i]) {
						c := rev
						c[j] = rune(i)
						next(i+1, c)
					}
				}
			} else {
				count++
			}
		}

		if count == l {
			for i, ch := range rev {
				if ch > 0 {
					dict[ch] = byte(i)
				}
			}

			if decypher(dict, rev) {
				res = true
			}
		}
	}
	next('A', [10]rune{})

	return res
}

func main() {
	// fmt.Println(isSolvable([]string{"AA", "BB"}, "DEC"))
	// fmt.Println(isSolvable([]string{"SEND", "MORE"}, "MONEY"))
	// fmt.Println(isSolvable([]string{"THIS", "IS", "TOO"}, "FUNNY"))
	fmt.Println(isSolvable([]string{"LEET", "CODE"}, "POINT"))

	fmt.Println(isSolvable([]string{"SEIS", "CATORCE", "SETENTA"}, "NOVENTA"))
}
