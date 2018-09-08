package main

import (
	"fmt"
	"math"
)

func repeatedDNASequences(s string) []string {
	nucleotides := map[rune]int{
		'A': 0,
		'C': 1,
		'G': 2,
		'T': 3,
	}

	modBase, max := int(math.Pow(4, 9)), -1
	indices := map[int]int{}
	count := map[int]int{}
	sum := 0

	runes := []rune(s)
	for i, ch := range runes {
		sum = sum % modBase
		sum *= 4
		sum += nucleotides[ch]

		if i >= 9 {
			if 1 == count[sum] {
				indices[sum] = i
			}
			count[sum]++
			if sum > max {
				max = sum
			}
		}

		//	fmt.Println(i, ch, sum)
	}

	//	fmt.Println(indices)

	result := []string{}
	for i := 0; i <= max; i++ {
		if count[i] > 1 {
			index := indices[i]
			result = append(result, s[index-9:index+1])
		}
	}
	return result
}

func main() {
	fmt.Println(repeatedDNASequences("AAAAAAAAAAA"))
	fmt.Println(repeatedDNASequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
}
