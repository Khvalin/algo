package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func solve(a []int) (int, int) {
	seen := map[int]bool{}

	sort.Ints(a)

	x, y := a[len(a)-1], 1

	yDivisors := []int{}

	for i := len(a) - 1; i >= 0; i-- {
		n := a[i]
		if seen[n] || x%n > 0 {
			yDivisors = append(yDivisors, n)
		}
		seen[n] = true
	}

	if len(yDivisors) > 0 {
		y = yDivisors[0]
	}

	return x, y
}

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}

func main() {
	n := 0

	fmt.Scan(&n)
	a := make([]int, n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := range a {
		scanner.Scan()
		a[i] = toInt(scanner.Bytes())
	}

	fmt.Println(solve(a))
}
