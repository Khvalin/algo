package main

import (
	"fmt"
	"sort"
)

//a+b, a+c, b+c, a+b+c
func solve(x ...int64) (int64, int64, int64) {
	var a, b, c int64
	/*
		b = x1 -a
		c = x2 - a
		a - 2a + x1 + x2 = x4
		a = -x4 +x1 + x2
	*/

	sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})

	a = -x[3] + x[0] + x[1]
	b = x[0] - a
	c = x[1] - a

	return a, b, c
}

func main() {
	var x1, x2, x3, x4 int64

	fmt.Scan(&x1, &x2, &x3, &x4)

	a, b, c := solve(x1, x2, x3, x4)
	fmt.Println(a, b, c)
}
