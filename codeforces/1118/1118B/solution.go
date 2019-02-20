package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(a []int) int {
	res := (0)

	a = append([]int{0}, a...)
	n := len(a)
	oddSum := make([]int, n)
	evenSum := make([]int, n)

	oddTotal := 0
	evenTotal := 0

	for i, w := range a {
		if i%2 == 0 {
			evenTotal += w
		} else {
			oddTotal += w
		}
		oddSum[i] = oddTotal
		evenSum[i] = evenTotal
	}

	for i := 1; i < n; i++ {
		even := evenSum[i-1] + oddTotal - oddSum[i]
		odd := oddSum[i-1] + evenTotal - evenSum[i]
		if even == odd {
			res++
		}
	}

	return res
}

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}

func main() {
	n := 0

	fmt.Scan(&n)
	a := make([]int, n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := range a {
		scanner.Scan()
		a[i] = toInt(scanner.Bytes())
	}

	fmt.Println(solve(a))
}
