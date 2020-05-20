package main

import "fmt"

func findAnagrams(s, p string) []int {
	type state struct {
		letters [28]int
		index   int
		count   int
	}
	pLetters := [28]int{}
	for _, ch := range p {
		pLetters[ch-'a']++
	}

	st := state{}
	res := []int{}

	for i, ch := range s {
		ord := ch - 'a'
		if pLetters[ord] == 0 {
			if st.count > 0 {
				// reset letters count
				for j := range st.letters {
					st.letters[j] = 0
				}
			}
			st.count = 0
			st.index = i + 1

			continue
		}

		st.letters[ord]++
		st.count++

		if st.count < len(p) && st.letters[ord] <= pLetters[ord] {
			// if there was no overflow, keep iterating
			continue
		}

		if st.count == len(p) && st.letters[ord] == pLetters[ord] {
			// match found
			res = append(res, st.index)
		}

		for st.count == len(p) || st.letters[ord] > pLetters[ord] {
			st.letters[s[st.index]-'a']--
			st.index++
			st.count--
		}
	}

	return res
}

func main() {
	fmt.Println(findAnagrams("abcab", "abc"))
}
