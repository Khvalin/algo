package main

import (
	"bufio"
	"fmt"
	"os"
)

// Range ...
type Range = struct {
	l   int
	r   int
	min uint32
}

var memo map[Range][]bool

func solve(a *[]uint32, l, r int, min uint32) []bool {
	if l > r {
		return nil
	}

	hash := Range{l, r, min}

	res, found := memo[hash]
	if found {
		return res
	}

	arr := *a

	res = make([]bool, 0, len(arr))

	for ok := true; ok && l <= r; {
		ok = true

		if arr[l] > min && (arr[l] == arr[r]) {
			left := solve(a, 1+l, r, arr[l])
			right := solve(a, l, r-1, arr[r])

			if len(left) > len(right) {
				res = append(res, false)
				res = append(res, left...)
				break
			}

			res = append(res, true)
			res = append(res, right...)

			break
		}

		if arr[l] > min && (arr[l] < arr[r] || arr[r] <= min) {
			res = append(res, false) //+ solve(a, 1+l, r, arr[l])
			min = arr[l]
			l++
			continue
		}

		if arr[r] > min {
			res = append(res, true) // + solve(a, l, r-1, arr[r])
			min = arr[r]
			r--
			continue
		}
		ok = false
	}

	memo[hash] = res
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
	memo = map[Range][]bool{}
	var n int

	fmt.Scan(&n)

	a := make([]uint32, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := range a {
		scanner.Scan()
		a[i] = toInt(scanner.Bytes())
	}

	ans := solve(&a, 0, len(a)-1, 0)

	fmt.Println(len(ans))
	for _, b := range ans {
		if b {
			fmt.Print("R")
		} else {
			fmt.Print("L")
		}
	}
	fmt.Println()
}
