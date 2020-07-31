package main

import "fmt"

func wordBreak(s string, wordDict []string) []string {
	type state struct {
		sIndex  int
		wIndex  int
		chIndex int
		chain   []int
	}

	var res []string

	if len(s) == 0 || len(wordDict) == 0 {
		return res
	}

	starts := map[byte][]int{}
	for j, w := range wordDict {
		if len(w) > 0 {
			starts[w[0]] = append(starts[w[0]], j)
		}
	}

	options := []state{}
	for _, i := range starts[s[0]] {
		options = append(options, state{sIndex: 1, wIndex: i, chIndex: 1})
	}

	for len(options) > 0 {
		opt := &options[0]
		//fmt.Printf("%d ",opt.sIndex)

		w := wordDict[opt.wIndex]

		if opt.sIndex == len(s) {
			if opt.chIndex == len(w) {
				r := ""
				for _, i := range opt.chain {
					r += wordDict[i] + " "
				}

				r += w
				res = append(res, r)
			}

			options = options[1:]
			continue
		}

		if opt.chIndex == len(w) {
			chain := append(opt.chain, opt.wIndex)

			var next []state
			for _, i := range starts[s[opt.sIndex]] {
				next = append(next, state{opt.sIndex + 1, i, 1, chain})
			}
			options = append(next, options[1:]...)

			continue
		}

		if w[opt.chIndex] != s[opt.sIndex] {
			options = options[1:]
			continue
		}

		opt.chIndex++
		opt.sIndex++
	}

	return res
}

func main() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}

	//	fmt.Printf("%s\n", wordBreak(s, wordDict))

	s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	wordDict = []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	fmt.Printf("%s\n", wordBreak(s, wordDict))
}
