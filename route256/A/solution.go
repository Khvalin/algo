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
	a, b, c, d := 0, 0, 0, 0

	fmt.Scan(&a, &b, &c, &d)
	days := map[int]int{}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < d; i++ {
		scanner.Scan()
		j := toInt(scanner.Bytes())

		scanner.Scan()
		days[j] = toInt(scanner.Bytes())
	}

	prev := -1
	output := bufio.NewWriter(os.Stdout)
	for i := a; i <= b; i++ {
		r := -1
		if i <= c {
			if n, found := days[i]; found && n != -1 {
				prev = n
			}

			r = prev
		}

		s := fmt.Sprintln(i, r)
		output.WriteString(s)
	}

	output.Flush()
}
