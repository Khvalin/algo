package main

import (
	"bufio"
	"fmt"
	"os"
)

type int = int64

func solve(h int, a []int) int {
	var r, sum int64
	//t := h
	for i, n := range a {
		sum += n

		//		t += n
		//1fmt.Println(t)

		if -sum >= h {
			return int64(i + 1)
		}
	}

	if sum >= 0 {
		return -1
	}

	c := h/-sum - 1
	r += c * int64(len(a))

	h += c * sum

	fmt.Println(h, c, r)

	for i := 0; h > 0; i++ {
		h += a[i]
		r++
	}

	return r
}

func toInt(buf []byte) int {
	n, mul := int64(0), int64(1)
	for i, v := range buf {
		if i == 0 && v == '-' {
			mul = -1
			continue
		}
		n = n*10 + int64(v-'0')
	}

	return mul * n
}

func main() {
	var h, r int64
	fmt.Scan(&h, &r)

	a := make([]int, r)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := range a {
		scanner.Scan()
		a[i] = toInt(scanner.Bytes())
	}

	ans := solve(h, a)
	fmt.Println(ans)
}
