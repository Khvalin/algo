package main

import "fmt"

func minInsertions(s string) int {
	memo := map[int]int{}

	var solve func(l, r int) int
	solve = func(l, r int) int {
		if l >= r {
			return 0
		}

		key := (l+1)*529 + r
		if res, f := memo[key]; f {
			return res
		}

		for l < r && s[l] == s[r] {
			l++
			r--
		}
		if l >= r {
			return 0
		}

		res, t := solve(l+1, r), solve(l, r-1)
		if t < res {
			res = t
		}
		res++
		memo[key] = res

		return res
	}

	return solve(0, len(s)-1)
}

func main() {
	t := "zjveiiwvc"
	r := minInsertions(t)
	fmt.Println(r)

}
