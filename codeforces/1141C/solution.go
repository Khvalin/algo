package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toInt(buf []byte) int {
	n, mul := 0, 1
	for i, v := range buf {
		if i == 0 && v == '-' {
			mul = -1
			continue
		}
		n = n*10 + int(v-'0')
	}

	return mul * n
}

func solve(n int, q []int) ([]int, bool) {
	p := []int{n}

	p = append(p, q...)

	min := n
	for i := 1; i < n; i++ {
		p[i] += p[i-1]
		if p[i] < min {
			min = p[i]
		}
	}

	fix := 1 - min
	used := map[int]bool{}

	for i := range p {
		p[i] += fix
		used[p[i]] = true
	}

	isValid := true
	for i := 1; isValid && (i <= n); i++ {
		isValid = used[i]
	}

	return p, isValid
}

func main() {
	n := 0

	fmt.Scan(&n)
	q := make([]int, n-1)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := range q {
		scanner.Scan()
		q[i] = toInt(scanner.Bytes())
	}

	p, isValid := (solve(n, q))
	if !isValid {
		p = []int{-1}
	}

	output := make([]string, len(p))

	for i, n := range p {
		output[i] = strconv.Itoa(n)
	}

	fmt.Println(strings.Join(output, " "))
}
