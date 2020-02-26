package main

import (
	"bufio"
	"fmt"
	"os"
	//	"sort"
)

func toInt(buf []byte) int {
	n := int(0)
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}

	return n
}

func readData() (n, m, k int, connected [][]int, special []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n = toInt(scanner.Bytes())

	scanner.Scan()
	m = toInt(scanner.Bytes())

	scanner.Scan()
	k = toInt(scanner.Bytes())

	special = make([]int, k)
	for i := range special {
		scanner.Scan()
		special[i] = toInt(scanner.Bytes())
	}
	// sort.Slice(special, func(i, j int) bool {
	// 	return special[i] < special[j]
	// })

	connected = make([][]int, 1+n)
	for i := int(0); i < m; i++ {
		scanner.Scan()
		x := toInt(scanner.Bytes())

		scanner.Scan()
		y := toInt(scanner.Bytes())

		connected[x] = append(connected[x], y)
		connected[y] = append(connected[y], x)
	}

	return
}

func solve(n, m, k int, connected [][]int, special []int) int {
	d, path := make([]int, n+1), make([]int, n+1)
	for i := range d {
		d[i] = m + 2
	}

	isSpecial := make([]bool, n+2)
	for _, n := range special {
		isSpecial[n] = true
	}

	q := []int{1}
	d[1] = 0

	for len(q) > 0 {
		p := q[len(q)-1]
		q = q[:len(q)-1]

		for _, n := range connected[p] {
			if d[n] > m {
				d[n] = d[p] + 1
				q = append(q, n)

				path[n] = p
			}
		}
	}

	res := 0

	c := 0
	last := -1
	for i := n; i > 0; i = path[i] {
		if isSpecial[i] {
			c++
			if last > 0 {
				t := d[i] + d[n] - d[last] + 1
				if t > res {
					res = t
				}
			}

			last = i
		}
	}

	if c < k {
		return d[n]
	}

	return res
}

func main() {
	fmt.Println(solve(readData()))
}
