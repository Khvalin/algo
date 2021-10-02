package main

import (
	"bufio"
	"fmt"
	"os"
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
	n := 0

	fmt.Scan(&n)
	a := make([]int, n)

	output := bufio.NewWriter(os.Stdout)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := range a {
		scanner.Scan()
		a[i] = toInt(scanner.Bytes())
		if i > 0 {
			a[i] += a[i-1]
		}
	}

	scanner.Scan()
	m := toInt(scanner.Bytes())

	for i := 0; i < m; i++ {
		scanner.Scan()
		q := toInt(scanner.Bytes())

		l, r := 0, n
		for l < r {
			mid := (l + r) >> 1
			if a[mid] < q {
				l = mid + 1
				continue
			}

			r = mid
		}

		output.WriteString(fmt.Sprintln(r + 1))
	}

	output.Flush()
}
