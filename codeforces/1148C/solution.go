package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(a []int) [][2]int {
	n := len(a)
	res := [][2]int{}

	diff := make([]int, n)

	for i, n := range a {
		diff[i] = (n - i + 1)

	}

	return res
}

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}

func main() {
	var n int

	fmt.Scan(&n)

	a := make([]int, n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := range a {
		scanner.Scan()
		a[i] = toInt(scanner.Bytes())
	}

	swaps := solve(a)
	for _, p := range swaps {
		fmt.Println(p[0], p[1])
	}
}
