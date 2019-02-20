package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(N int, a []int) (string, [][]int) {
	const NO = "NO"
	const YES = "YES"

	res := [][]int{}
	for i := 0; i < N; i++ {
		res = append(res, make([]int, N))
	}

	counts := map[int]int{}
	for _, n := range a {
		counts[n]++
	}

	oddCount := 0
	centerValue := 0
	for k, v := range counts {
		if v%2 == 1 {
			oddCount++
			centerValue = k
		}
	}

	if oddCount > 1 {
		return NO, res
	}

	mid := N >> 1

	if oddCount == 1 {
		if N%2 == 1 {
			res[mid][mid] = centerValue
			counts[centerValue]--
		} else {
			return NO, res
		}
	}

	center := N>>1 + N%2

	i := 0
	for len(counts) > 0 {
		n := 0

		for k, v := range counts {
			if v > 3 {
				n = k
				counts[k] -= 4
			}
			if counts[k] < 4 {
				delete(counts, k)
			}
			if n > 0 {
				break
			}
		}

		if n <= 0 {
			return NO, res
		}

		x, y := i/center, i%center

		res[x][y] = n

		res[N-1-x][y] = n
		res[x][N-1-y] = n
		res[N-1-x][N-1-y] = n

		i++
	}

	return YES, res
}

func toInt(buf []byte) int {
	n := 0
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return n
}

func main() {
	n := 0

	fmt.Scan(&n)
	a := make([]int, n*n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := range a {
		scanner.Scan()
		a[i] = toInt(scanner.Bytes())
	}

	fmt.Println(solve(n, a))
}
