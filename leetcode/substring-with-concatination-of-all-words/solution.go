package main

import (
	"fmt"
)

func findSubstring(s string, words []string) []int {
	result := []int{}
	LEN := 0
	if len(words) > 0 {
		LEN = len(words[0])
	}

	resetWordsMap := func(res map[string]int, words []string) map[string]int {
		for i := 0; i < len(words); i++ {
			res[words[i]] = 0
		}

		for i := 0; i < len(words); i++ {
			res[words[i]]++
		}

		return res
	}

	count := 0
	start := -1

	wordsMap := map[string]int{}
	resetWordsMap(wordsMap, words)

	for i := 0; i <= len(s)-LEN; {
		w := s[i : i+LEN]

		val, ok := wordsMap[w]
		if ok && val > 0 {
			if 0 == count {
				start = i
			}
			wordsMap[w]--
			count++
			//		fmt.Println(i, wordsMap)
			if len(words) == count {
				result = append(result, start)
				// fall through to reset counters
			} else {
				i += LEN
				continue
			}
		}

		if count > 0 {
			resetWordsMap(wordsMap, words)
			count = 0
			i = start
		}

		i++
	}

	return result
}

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))

}
