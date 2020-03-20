package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(n uint64) string {
	if n == 1 {
		return "-1"
	}

	sum := 0
	r := make([]byte, n)
	for i := range r {
		d := 2
		for d < 9 && (sum+d)%9 == 0 {
			d *= 2
		}
		if d > 9 {
			return "-1"
		}
		sum += d
		r[i] = '0' + byte(d)
	}
	r[n-1] = '9'

	return string(r)
}

func main() {
	toInt := func(buf []byte) uint64 {
		var n uint64
		for _, v := range buf {
			n = n*10 + uint64(v-'0')
		}
		return n
	}

	var n uint64

	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := uint64(0); i < n; i++ {
		scanner.Scan()
		input := toInt(scanner.Bytes())
		fmt.Println(solve(input))
	}

}
