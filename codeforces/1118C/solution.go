package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(N int, a []int) (bool, [][]int) {
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
		return false, nil
	}

	mid := N >> 1

	if oddCount == 1 {
		if N%2 == 1 {
			res[mid][mid] = centerValue
			counts[centerValue]--
		} else {
			return false, nil
		}
	}

	center := N >> 1

	components := [][][]int{}

	x, y := 0, 0
	for x < center && y < center {
		com := [][]int{
			[]int{x, y},
			[]int{N - 1 - x, y},
			[]int{x, N - 1 - y},
			[]int{N - 1 - x, N - 1 - y},
		}

		components = append(components, com)

		y++
		if y >= center {
			y = 0
			x++
		}
	}

	if N%2 == 1 {
		x, y = center, 0
		for y < center {
			com := [][]int{
				[]int{x, y},
				[]int{x, N - 1 - y},
			}

			components = append(components, com)

			com = [][]int{
				[]int{y, x},
				[]int{N - 1 - y, x},
			}

			components = append(components, com)

			y++
		}
	}

	for _, com := range components {
		slotsCount := len(com)
		n := 0

		for k, v := range counts {
			if v >= slotsCount {
				n = k
				counts[k] -= slotsCount
			}
			if counts[k] == 0 {
				delete(counts, k)
			}
			if n > 0 {
				break
			}
		}

		if n <= 0 {
			return false, nil
		}

		for _, pair := range com {
			res[pair[0]][pair[1]] = n
		}
	}

	return true, res
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

	valid, res := (solve(n, a))
	if !valid {
		fmt.Println("NO")
		return
	}

	fmt.Println("YES")
	for _, r := range res {
		for _, n := range r {
			fmt.Print(n, " ")
		}
		fmt.Println()
	}
}
