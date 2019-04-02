package main

import (
	"bufio"
	"fmt"
	"os"
)

type int = int64

type breakPoint struct {
	hours int
	sum   int
}

func solve(h int, a []int32) int {
	var r, sum int64

	L := len(a)

	j, t, k := -1, int64(0), -1
	bpts := make([]breakPoint, L+1)

	for i, n := range a {
		sum += int64(n)

		t += int64(n)

		if t < 0 || i == L-1 {
			k++
			bpts[k] = breakPoint{
				hours: int64(i - j),
				sum:   t,
			}

			t = 0
			j = i
		}

		if -sum >= h {
			return int64(i + 1)
		}
	}

	if sum >= 0 {
		return -1
	}

	bpts = bpts[:k+1]

	c := h / -sum

	r += c * int64(len(a))

	ans := r

	winCondition := func(h int) (bool, int) {
		var r int64

		for _, point := range bpts {
			if h <= 0 {
				return true, r
			}
			r += point.hours
			h += point.sum
		}

		return h <= 0, r
	}

	left, right := int64(0), r-1

	for left < right-1 {
		//	fmt.Println(left, right)
		mid := (left + right) >> 1
		condition, res := winCondition(h + sum*mid)

		if condition {
			ans = res + mid*int64(len(a))
			right = mid
		} else {
			left = mid
		}
	}

	return ans
}

func toInt(buf []byte) int32 {
	n, mul := int32(0), int32(1)
	for i, v := range buf {
		if i == 0 && v == '-' {
			mul = -1
			continue
		}
		n = n*10 + int32(v-'0')
	}

	return mul * n
}

func main() {
	var h, r int64
	fmt.Scan(&h, &r)

	a := make([]int32, r)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := range a {
		scanner.Scan()
		a[i] = toInt(scanner.Bytes())
	}

	ans := solve(h, a)
	fmt.Println(ans)
}
