package main

import "fmt"

func longestSubstring(s string, k int) int {
	count := make([][28]int, len(s)+2)

	for i, ch := range s {
		ord := ch - 'a'

		if i > 0 {
			count[i+1] = count[i]
		}

		count[i+1][ord]++
	}

	var solve func(start, end int) int
	solve = func(start, end int) int {
		//	fmt.Println(start, end)
		res := 0
		i := start
		for j := i; j < end; j++ {
			ord := s[j] - 'a'
			if count[end][ord]-count[i][ord] < k {
				a, b := solve(i, j), solve(j+1, end)
				if a > b {
					return a
				}

				return b
			}

			f := true
			c := count[i]
			for l, n := range count[j+1] {
				f = f && (n == c[l] || n-c[l] >= k)
			}

			if f && res < j-i+1 {
				res = j - i + 1
			}
		}

		return res
	}

	return solve(0, len(s))
}

func main() {
	fmt.Println(longestSubstring("aaabb", 3))
	fmt.Println(longestSubstring("aaaaaaaaazcccccddddd", 5))
	fmt.Println(longestSubstring("bbaaacbd", 3))
}
