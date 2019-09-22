package main

import (
	"bufio"
	"fmt"
	"os"
)

type state struct {
	index int
	prev  uint32
}

func solve(a, b []uint32) float64 {
	m := map[state]float64{}

	var f func(index int, prev uint32) float64
	f = func(index int, prev uint32) float64 {
		if index > len(a)-1 {
			return 0
		}

		s := state{index, prev}
		res, has := m[s]
		if has {
			return res
		}

		r := 10e18 + 1
		for h := uint32(0); h < 3; h++ {
			t := r
			if a[index]+h != prev {
				t = float64(h*b[index]) + (f(index+1, a[index]+h))
			}

			if t < r {
				r = t
			}
		}

		m[s] = r

		return r
	}

	res := f(0, a[0]-1)
	//fmt.Println(m)

	return res
}

func toInt(buf []byte) uint32 {
	n := uint32(0)
	for _, v := range buf {
		n = n*10 + uint32(v-'0')
	}

	return n
}

func main() {
	var q uint32

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	q = toInt(scanner.Bytes())

	for i := uint32(0); i < q; i++ {
		var n int

		scanner.Scan()
		n = int(toInt(scanner.Bytes()))

		a := make([]uint32, n)
		b := make([]uint32, n)
		for j := range a {
			scanner.Scan()
			a[j] = toInt(scanner.Bytes())

			scanner.Scan()
			b[j] = toInt(scanner.Bytes())
		}

		fmt.Printf("%.0f\n", solve(a, b))
	}
}
