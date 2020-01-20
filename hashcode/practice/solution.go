package main

import (
	"bufio"
	"fmt"
	"os"
)

func toInt(buf []byte) uint {
	n := uint(0)
	for _, v := range buf {
		n = n*10 + uint(v-'0')
	}

	return n
}

func solve(m, n uint, s []uint) []uint {
	type state struct {
		total   uint
		indices []uint
	}
	memo := make([]*state, n+1)

	var solve func(ind uint) state
	solve = func(ind uint) state {
		if ind >= n || s[ind] > m {
			return state{}
		}

		r := memo[ind]
		if r != nil {
			return *r
		}

		var st state
		max := uint(0)
		for i := ind; i < n; i++ {
			opt := solve(i + 1)
			sum := s[ind] + opt.total
			if sum <= m && sum > max {
				max = sum
				st = opt
			}
		}

		if max > 0 {
			st.total = max
			st.indices = append([]uint{ind}, st.indices...)
		}

		memo[ind] = &st

		return st
	}

	r := solve(0)

	return r.indices
}

func main() {
	var n uint
	var m uint

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	m = toInt(scanner.Bytes())

	scanner.Scan()
	n = toInt(scanner.Bytes())

	s := make([]uint, n)
	for i := uint(0); i < n; i++ {
		scanner.Scan()
		s[i] = toInt(scanner.Bytes())
	}

	res := solve(m, n, s)
	fmt.Println(len(res))
	for i := range res {
		prefix := ""
		if i > 0 {
			prefix = " "
		}
		fmt.Print(prefix, res[i])
	}

	fmt.Println()
}
