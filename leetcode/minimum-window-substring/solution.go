package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	toVector := func(s string) map[byte]int {
		vector := map[byte]int{}
		for i := 0; i < len(s); i++ {
			_, found := vector[s[i]]
			if !found {
				vector[s[i]] = 0
			}

			vector[s[i]]++
		}

		return vector
	}

	tVector := toVector(t)
	sVector := map[byte]int{}

	//fmt.Println(tVector)
	//fmt.Println(sVector)

	result := ""
	minLen := math.MaxInt32
	foundCount := 0
	var matchIndices []int

	for i := 0; i < len(s); i++ {
		ch := s[i]
		count, found := tVector[ch]

		if found {
			_, found := sVector[ch]
			if !found {
				sVector[ch] = 0
			}

			if sVector[ch] >= count {
				for i := 0; i < len(matchIndices); i++ {
					if s[matchIndices[i]] == ch {
						matchIndices = matchIndices[i:]
						break
					}
				}
			} else {
				sVector[ch]++
			}

			//fmt.Println(sVector)

			matchIndices = append(matchIndices, i)

			if sVector[ch] == count {
				foundCount++
				//fmt.Println(string(ch), foundCount)

				if foundCount == len(tVector) {
					var start int
					start, matchIndices = matchIndices[0], matchIndices[1:]
					sVector[s[start]]--
					curLen := i - start + 1
					if curLen < minLen {
						result = s[start : i+1]
						//fmt.Println(result)
						minLen = curLen
					}

					//fmt.Println(minLen, sVector)
					foundCount--
				}
			}

		}
	}

	return result
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("ABBA", "AA"))
	fmt.Println(minWindow("ABDDCAABC", "ABC"))
	fmt.Println(minWindow("CABBANC", "ABC"))
	fmt.Println(minWindow("HELLOW", "Z"))
}
