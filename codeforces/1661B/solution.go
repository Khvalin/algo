package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(a []uint64) []uint64 {
	getPow := func(d uint64) uint64 {
		i := uint64(0)
		for d%2 == 0 && d > 1 {
			d = d >> 1
			i++
		}

		return i
	}

	getMin := func(n uint64) uint64 {
		max := 15 - getPow(n)
		res := max

		for i := uint64(0); i <= max; i++ {
			t := (n + i) % 32768
			r := i
			if t > 0 {
				r += 15 - getPow(t)
			}
			if r < res {
				res = r
			}
		}

		return uint64(res)
	}

	for i, n := range a {
		a[i] = getMin(n % 32768)
	}

	return a
}

func main() {
	toInt := func(buf []byte) uint64 {
		var n uint64
		for _, v := range buf {
			n = n*10 + uint64(v-'0')
		}
		return n
	}

	var n uint64

	fmt.Scan(&n)
	b := make([]uint64, n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := range b {
		scanner.Scan()
		b[i] = toInt(scanner.Bytes())
	}

	a := solve(b)
	for _, x := range a {
		fmt.Print(x, " ")
	}
	fmt.Println()
}
