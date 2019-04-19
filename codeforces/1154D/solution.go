package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(n int, b, a int64, s []bool) int {
	var (
		i int
	)

	bCur, aCur := b, a
	for ; i < n; i++ {
		if aCur <= 0 && bCur <= 0 {
			break
		}

		f := s[i]
		if f && bCur > 0 && aCur < a {
			aCur++
			bCur--
		} else {
			if aCur > 0 {
				aCur--
			} else {
				bCur--
			}
		}
	}

	return i
}

func main() {
	var (
		n    int
		b, a int64
		s    []bool
	)

	fmt.Scan(&n, &b, &a)

	s = make([]bool, n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := range s {
		scanner.Scan()
		s[i] = '1' == (scanner.Bytes())[0]
	}

	r := solve(n, b, a, s)
	fmt.Println(r)
}
