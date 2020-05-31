package main

import "fmt"

// https://en.wikipedia.org/wiki/Wagner%E2%80%93Fischer_algorithm
func minDistance(word1, word2 string) int {
	min := func(nums ...int) int {
		res := nums[0]
		for _, n := range nums {
			if n < res {
				res = n
			}
		}

		return res
	}

	m, n := len(word1)+1, len(word2)+1
	d := make([][]int, m)
	for i := range d {
		d[i] = make([]int, n)
	}

	// source prefixes can be transformed into empty string by dropping all characters
	for i := 0; i < m; i++ {
		d[i][0] = i
	}

	// target prefixes can be reached from empty source prefix by inserting every character
	for j := 0; j < n; j++ {
		d[0][j] = j
	}

	for j := 1; j < n; j++ {
		for i := 1; i < m; i++ {
			substitutionCost := 1
			if word1[i-1] == word2[j-1] {
				substitutionCost = 0
			}

			d[i][j] = min(d[i-1][j]+1, // deletion
				d[i][j-1]+1,                  // insertion
				d[i-1][j-1]+substitutionCost) // substitution
		}
	}

	return d[m-1][n-1]
}

func main() {
	fmt.Println(minDistance("horse", "ros"))
	fmt.Print(minDistance("intention", "execution"))
}
