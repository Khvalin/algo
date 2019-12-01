package main

import "fmt"

func palindromePartition(str string, k int) int {
	const nmax = 110
	memo := map[uint]int{}

	var calc func(s, e int) int
	calc = func(s, e int) int {
		if s > e {
			return 65536
		}

		key := uint(s*nmax + e)
		r, f := memo[key]
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

		memo[key] = r

		return r
	}

	l := len(str)
	p := [nmax]int{}
	for i := 0; i < k; i++ {
		p[i] = 1
	}
	m := l - k + 1
	p[0] = m

	perms := [][nmax]int{p}

	res := 65536

	for len(perms) > 0 {
		p := &perms[0]
		/* 		for i := 0; i < k+4; i++ {
		   			fmt.Print(p[i], " ")
		   		}
		   			fmt.Println(res) */

		j := p[k+2]

		if p[0] < 1 || p[k] > 0 {
			perms = perms[1:]
		}

		if p[0] < 1 || p[k] > 0 {
			continue
		}

		s := 0
		l := 0
		for i := 0; i < k; i++ {
			n := p[i]
			s += calc(l, l+n-1)
			l += n
		}

		if s < res {
			res = s
		}

		valid := true
		for j < k-1 && p[j] == 1 {
			valid = false
			j++
		}
		if j >= k-1 {
			j = 0
		}

		p[k+2] = j

		p[j]--
		j++
		p[j]++

		nextP := [nmax]int{}

		valid = valid && p[j] < m && j < k-1 && p[j] > 1

		for i := 0; i < k+3; i++ {
			if p[i] < 1 && i < k {
				valid = false
				break
			}
			nextP[i] = p[i]
		}

		if valid {
			nextP[k+2] = j + 1
			perms = append(perms, nextP)
		}

	}

	return res
}

func main() {
	r := 0
	r = palindromePartition("oiwwhqjkb", 1)
	fmt.Println(r)

	r = palindromePartition("abc", 3)
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
