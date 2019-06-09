package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(a, b []int, ta, tb, k int) int {
	for len(b) > 0 && b[0] < a[0]+ta {
		b = b[1:]
	}

	if (len(a)) <= k || (len(b)) <= k {
		return -1
	}

	if b[len(b)-k] < a[len(a)-1] {
		return -1
	}

	for i := 0; i < k; i++ {
		if a[1]-a[0] < b[1]-b[0] {
			b = b[1:]
		} else {
			a = a[1:]
		}

		for b[0] < a[0]+ta {
			b = b[1:]
		}
	}

	return b[0] + tb
}

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}

func main() {
	var n, m, ta, tb, k int

	fmt.Scan(&n, &m, &ta, &tb, &k)

	a, b := make([]int, n), make([]int, m)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := range a {
		scanner.Scan()
		a[i] = toInt(scanner.Bytes())
	}

	for i := range b {
		scanner.Scan()
		b[i] = toInt(scanner.Bytes())
	}

	r := solve(a, b, ta, tb, k)
	fmt.Println(r)
}
