package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type UInt64Slice []int64

func (p UInt64Slice) Len() int           { return len(p) }
func (p UInt64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p UInt64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func toInt(buf []byte) int64 {
	var n int64
	for _, v := range buf {
		n = n*10 + int64(v-'0')
	}
	return n
}

func readInput() (UInt64Slice, int64) {
	var k int64

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

func solve(a UInt64Slice, k int64) int64 {
	n := len(a)

	for k > 0 {
		if a[n-1] == a[0] {
			break
		}

		l := int64(sort.Search(n, func(i int) bool {
			return a[i] > a[0]
		}))

		r := int64(sort.Search(n, func(i int) bool {
			return a[n-i-1] < a[len(a)-1]
		}))

		dl := a[l] - a[0]
		if dl*l > k {
			dl = (dl * l) / k
		}

		dr := a[n-1] - a[n-int(r)-1]
		if r*dr > k {
			dr = (r * dr) / k
		}

		/* 		fmt.Println(a)
		   		fmt.Println(l, dl)
		   		fmt.Println(r, dr) */

		if dl*l < r*dr {
			for i := l - 1; k > 0 && i >= 0; i-- {
				if k < dl {
					dl = (k)
				}
				a[i] += dl
				k -= dl
			}
		} else {
			r = int64(n) - r

			for i := r; i < int64(n) && k > 0; i++ {
				if k < dr {
					dr = (k)
				}

				a[i] -= dr
				k -= dr
			}
		}
		/* 		fmt.Println(a) */
	}

	return a[n-1] - a[0]
}

func main() {
	a, k := readInput()
	sort.Sort(a)

	r := solve(a, (k))

	fmt.Println(r)
}
