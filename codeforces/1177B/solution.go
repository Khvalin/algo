package main

import (
	"fmt"
)

/*
1 1
10 11
100 192
1000 2893
10000 38894
100000 488895
1000000 5888896
10000000 68888897
100000000 788888898
1000000000 8888888899

*/

// Solve returns k-th char from 123456789101112131415161718192021222324252627282930313233343536
func Solve(n uint64) byte {
	pows := [15]uint64{1}

	for i := 1; i <= 12; i++ {
		pows[i] = pows[i-1] * 10
	}

	i, l := uint64(0), uint64(0)
	p := uint64(0)

	for next := uint64(1); l+next <= n; next = 1 + (pows[i]-p)*i {
		l += next
		//fmt.Println(l)

		p = pows[i]
		i++
	}

	l -= i - 1

	t := pows[i-1] + (n-l)/i

	s := fmt.Sprint(t)
	ind := (n - l) % i

	//	fmt.Println(l, pows[i-1], s, ind)

	return s[ind] - '0'
}

func main() {
	var n uint64

	fmt.Scanf("%d", &n)

	ans := Solve(n)
	fmt.Println(ans)
}
