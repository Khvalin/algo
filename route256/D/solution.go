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

	pr := map[int]map[int]int{}

	output := bufio.NewWriter(os.Stdout)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		scanner.Scan()
		w := toInt(scanner.Bytes())

		scanner.Scan()
		c := toInt(scanner.Bytes())

		for j := 0; j < c; j++ {
			scanner.Scan()
			id := toInt(scanner.Bytes())

			scanner.Scan()
			d := toInt(scanner.Bytes())

			if pr[id] == nil {
				pr[id] = map[int]int{}
			}

			pr[id][w] = d
		}
	}

	scanner.Scan()
	m := toInt(scanner.Bytes())
	for i := 0; i < m; i++ {
		scanner.Scan()
		id := toInt(scanner.Bytes())

		scanner.Scan()
		p := toInt(scanner.Bytes())

		f := false
		res := [2]int{-1, -1}
		if units, found := pr[id]; found {
			for k, v := range units {
				if !f || v < res[1] || v == res[1] && k < res[0] {
					res[0] = k
					res[1] = v
				}

				f = true
			}

			res[1] += p
		}

		output.WriteString(fmt.Sprintln(res[0], res[1]))
	}

	output.Flush()
}
