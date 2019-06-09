package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(a []uint32) (int, int) {
	const N = uint32(1e7 + 10)

	min := uint64(1<<64 - 1)

	var (
		ans1 int
		ans2 int
	)

	m := map[uint32]int{}
	for i, n := range a {
		m[n] = 1 + i
	}

	for gcd := uint32(1); gcd < N; gcd++ {
		index1, index2 := -1, -1

		var (
			first  uint32
			second uint32
		)

		for j := gcd; j < N; j += gcd {
			i := m[j]
			if i > 0 {
				if index1 > 0 {
					second = j
					index2 = i

					break
				} else {
					first = j
					index1 = i
				}
			}
		}

		if first > 0 && second > 0 {
			gres := uint64(first*second) / uint64(gcd)

			if gres < min {
				min = gres
				ans1, ans2 = index1, index2
			}
		}
	}

	if ans1 > ans2 {
		ans1, ans2 = ans2, ans1
	}

	return ans1, ans2
}

func toInt(buf []byte) uint32 {
	n := uint32(0)
	for _, v := range buf {
		n = n*10 + uint32(v-'0')
	}

	return n
}

func main() {
	var n int

	fmt.Scan(&n)

	a := make([]uint32, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := range a {
		scanner.Scan()
		a[i] = toInt(scanner.Bytes())
	}

	i, j := solve(a)
	fmt.Println(i, j)
}
