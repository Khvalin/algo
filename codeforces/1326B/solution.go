package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(b []int64) []int64 {
	var max int64
	for i := range b {
		b[i] += max
		if b[i] > max {
			max = b[i]
		}
	}
	return b
}

func main() {
	toInt := func(buf []byte) int64 {
		var n int64
		var d int64 = 1
		for _, v := range buf {
			if v == '-' {
				d = -1
				continue
			}
			n = n*10 + int64(v-'0')
		}
		return d * n
	}

	var n int64

	fmt.Scan(&n)
	b := make([]int64, n)

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
