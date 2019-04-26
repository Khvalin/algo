package main

import (
	"bufio"
	"fmt"
	"os"
)

func toInt(buf []byte) int {
	n, mul := int(0), int(1)
	for i, v := range buf {
		if i == 0 && v == '-' {
			mul = int(-1)
			continue
		}
		n = n*10 + int(v-'0')
	}

	return mul * n
}

func solve(a []int) (int, int) {
	fmt.Println(a)
	return 0, 0
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

	i, j := solve(a)
	fmt.Println(i, j)
}
