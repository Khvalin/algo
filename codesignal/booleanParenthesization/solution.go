package main

import "fmt"

func booleanParenthesization(expression string) int {
	var mod uint = 1003
	nmax := 1000

	memo := make([][2]uint, nmax*nmax+nmax)

	var count func(l, r int) [2]uint
	count = func(l, r int) [2]uint {
		key := l*1e3 + r
		res := memo[key]

		if res[0] > 0 || res[1] > 0 {
			return res
		}

		if l >= r {
			return res
		}

		if r-l == 1 {
			if expression[l] == 'T' {
				res[1]++
			} else {
				res[0]++
			}

			return res
		}

		for i := l + 1; i < r; i += 2 {
			lcount, rcount := count(l, i), count(i+1, r)
			switch expression[i] {
			case '&':
				res[0] += lcount[0]*rcount[0] + lcount[0]*rcount[1] + lcount[1]*rcount[0]
				res[1] += rcount[1] * lcount[1]
				break
			case '|':
				res[0] += lcount[0] * rcount[0]
				res[1] += lcount[0]*rcount[1] + lcount[1]*rcount[0] + lcount[1]*rcount[1]
				break
			case '^':
				res[0] += rcount[0]*lcount[0] + rcount[1]*lcount[1]
				res[1] += rcount[0]*lcount[1] + rcount[1]*lcount[0]
				break
			}

			res[0] %= mod
			res[1] %= mod
		}

		memo[key] = res

		return res
	}

	res := count(0, len(expression))

	return int(res[1] % mod)
}

func main() {
	//fmt.Println( booleanParenthesization("F|T^F"))
	fmt.Println(5, booleanParenthesization("T&T|F^F"))
	fmt.Println(161, booleanParenthesization("T^T|T^F&F^T|F^T|F^F|T^F^F|T&T&T^F|T&F|F^T^F&T^F"))
}
