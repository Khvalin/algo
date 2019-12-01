package main

import "fmt"

func palindromePartition(str string, k int) int {
	const nmax = 110
	wordMemo := map[uint]int{}

	var calc func(s, e int) int
	calc = func(s, e int) int {
		if s > e {
			return nmax
		}

		key := uint(s*nmax + e)
		r, f := wordMemo[key]
		if f {
			return r
		}

		r = 0
		iMax := (e - s) >> 1
		for i := 0; i <= iMax; i++ {
			if str[s+i] != str[e-i] {
				r++
			}
		}

		wordMemo[key] = r

		return r
	}

	memo := map[uint]int{}
	var solve func(start, k int) int
	solve = func(start, k int) int {
		res := nmax
		if start >= len(str) {
			if k == 0 {
				return 0
			}

			return res
		}

		if k < 0 {
			return res
		}

		key := uint(start*nmax + k)
		r, f := memo[key]
		if f {
			return r
		}

		s := start
		if k == 1 {
			s = len(str) - 1
		}

		for i := s; i < len(str); i++ {
			t := calc(start, i) + solve(i+1, k-1)
			if t < res {
				res = t
			}
		}

		memo[key] = res

		return res
	}

	return solve(0, k)
}

func main() {
	r := 0

	r = palindromePartition("abc", 2)
	fmt.Println(1, r)

	r = palindromePartition("oiwwhqjkb", 1)
	fmt.Println(r)

	r = palindromePartition("aabbc", 3)
	fmt.Println(r)

	r = palindromePartition("leetcode", 8)
	fmt.Println(r)

	r = palindromePartition("mepekjkpgihfcg", 12)
	fmt.Println(r)

	r = palindromePartition("qiuihduhbwmxqlqlcfhpfxduz", 4)
	fmt.Println(8, r)

	r = palindromePartition("faaglagedtwnejzpuarkgwgoefwra", 27)
	fmt.Println(0, r)
}
