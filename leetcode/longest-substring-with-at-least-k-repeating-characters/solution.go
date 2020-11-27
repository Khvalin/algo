package main

import "fmt"

func longestSubstring(s string, k int) int {
	res := 0
	count := make([][28]int, len(s)+2)

	for i, ch := range s {
		ord := ch - 'a'

		if i > 0 {
			count[i+1] = count[i]
		}

		count[i+1][ord]++
	}

	for i := 0; i <= len(s)-k; i++ {
		c := count[i]
		for j := i + k; j <= len(s); j++ {

			f := true
			for l, n := range count[j] {
				f = f && (n-c[l] >= k || n == c[l])
			}

			if !f {
				continue
			}

			if j-i > res {
				res = j - i
			}
		}
	}

	return res
}

func main() {
	fmt.Println(longestSubstring("aaaaaaaaazcccccddddd", 5))
}
