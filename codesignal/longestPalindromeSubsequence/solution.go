package longestPalindromeSubsequence

import (
	"strings"
)

func longestPalindromeSubsequence(input string) string {
	str := strings.Replace(input, " ", "", -1)

	memo := map[int]string{}

	var findLongestPalindrome func(left, right int) string
	findLongestPalindrome = func(left, right int) string {
		if left > right {
			return ""
		}
		key := left*(len(str)+1) + right
		r, found := memo[key]
		if found {
			return r
		}

		ch := string(str[left])

		for i := right; i >= left; i-- {
			if str[i] == str[left] {
				next := ""
				if i == left {
					next = ch
				} else {
					next = ch + findLongestPalindrome(left+1, i-1) + ch
				}
				if len(next) >= len(r) {
					r = next
				}
			}
		}

		next := findLongestPalindrome(left+1, right)
		if len(next) > len(r) {
			r = next
		}

		memo[key] = r
		return r
	}

	return findLongestPalindrome(0, len(str)-1)
}
