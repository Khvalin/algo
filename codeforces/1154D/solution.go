package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(n, b, a int64, s []bool) int {
	var (
		i int
		f bool
	)

	//	bCur, aCur := b, a
	for i, f = range s {
		if f {

		}
	}

	return i
}

func main() {
	var (
		n, b, a int64
		s       []bool
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
