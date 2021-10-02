package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func toInt(buf []byte) int {
	n := 0
	d := 1

	for i, v := range buf {
		if i == 0 && v == '-' {
			d = -1
			continue
		}
		n = n*10 + int(v-'0')
	}
	return d * n
}

func main() {
	m, n := 0, 0

	fmt.Scan(&m, &n)
	a := make([]int, m)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := range a {
		scanner.Scan()
		a[i] = toInt(scanner.Bytes())
	}

	sort.Ints(a)

	res := -1
	j := 0
	prev := a[0] - 1
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == prev {
			continue
		}

		j++
		if j == n {
			res = a[i]
			break
		}

		prev = a[i]
	}

	fmt.Println(res)
}
