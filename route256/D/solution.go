package main

import (
	"bufio"
	"fmt"
	"os"
)

const UintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64

const (
	MaxInt  = 1<<(UintSize-1) - 1 // 1<<31 - 1 or 1<<63 - 1
	MinInt  = -MaxInt - 1         // -1 << 31 or -1 << 63
	MaxUint = 1<<UintSize - 1     // 1<<32 - 1 or 1<<64 - 1
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

	pr := map[int][]int{}

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
				pr[id] = make([]int, 2)
				pr[id][0] = w
				pr[id][1] = d
			}

			if pr[id][1] > d || pr[id][1] == d && pr[id][0] > w {
				pr[id][0] = w
				pr[id][1] = d
			}
		}
	}

	scanner.Scan()
	m := toInt(scanner.Bytes())
	for i := 0; i < m; i++ {
		scanner.Scan()
		id := toInt(scanner.Bytes())

		scanner.Scan()
		p := toInt(scanner.Bytes())

		res := [2]int{-1, -1}
		if units, found := pr[id]; found {
			res[0] = units[0]
			res[1] = units[1] + p
		}

		output.WriteString(fmt.Sprintln(res[0], res[1]))
	}

	output.Flush()
}
