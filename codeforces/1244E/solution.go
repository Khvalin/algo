package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type UInt64Slice []uint64

func (p UInt64Slice) Len() int           { return len(p) }
func (p UInt64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p UInt64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func toInt(buf []byte) uint64 {
	var n uint64
	for _, v := range buf {
		n = n*10 + uint64(v-'0')
	}
	return n
}

func readInput() (UInt64Slice, uint64) {
	var k uint64

	n := 0

	fmt.Scan(&n, &k)
	a := make(UInt64Slice, n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := range a {
		scanner.Scan()
		a[i] = toInt(scanner.Bytes())
	}

	return a, k
}

func solve(a UInt64Slice, k int64) uint64 {
	n := len(a)

	for k > 0 {
		fmt.Println(a, k)

		if a[n-1] == a[0] {
			break
		}

		l := sort.Search(n, func(i int) bool {
			return a[i] > a[0]
		})

		r := sort.Search(n, func(i int) bool {
			return a[n-i-1] < a[len(a)-1]
		})

		if l <= r {
			d := a[l] - a[0]
			dk := int64(d)
			for i := l - 1; k > 0 && i >= 0; i-- {
				if k >= dk {
					a[i] += d
					k -= dk
				} else {
					k = 0
				}

			}
			continue
		}

		r = n - r
		d := a[n-1] - a[r-1]
		dk := int64(d)

		for i := r; i < n && k > 0; i++ {
			if k < dk {
				d = uint64(k)
				dk = k
			}

			a[i] -= d
			k -= dk
		}
	}

	return a[n-1] - a[0]
}

func main() {
	a, k := readInput()
	sort.Sort(a)

	r := solve(a, int64(k))

	fmt.Println(r)
}
