package main

import (
	"bufio"
	"fmt"
	"os"
)

func touint64(buf []byte) uint64 {
	n := uint64(0)
	for _, v := range buf {
		n = n*10 + uint64(v-'0')
	}

	return n
}

func readData() []uint64 {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n := touint64(scanner.Bytes())

	r := make([]uint64, n)
	for i := range r {
		scanner.Scan()
		r[i] = touint64(scanner.Bytes())
	}

	return r
}

func solve(heights []uint64) []uint64 {
	L := len(heights)
	res, t := make([]uint64, L), make([]uint64, L)
	visited := make([]map[uint64]bool, L)

	var max uint64

	for i, h := range heights {
		if visited[i] == nil {
			visited[i] = map[uint64]bool{}
		}

		if visited[i][h] {
			continue
		}
		visited[i][h] = true

		copy(t, heights)

		s := h

		for j := i - 1; j >= 0; j-- {
			if t[j] > t[j+1] {
				t[j] = t[j+1]
			}
			s += t[j]
		}

		for j := i + 1; j < L; j++ {
			if t[j] > t[j-1] {
				t[j] = t[j-1]
			}
			s += t[j]
		}

		if s > max {
			max = s
			copy(res, t)
		}
	}

	return res
}

func main() {
	heights := readData()
	heights = solve(heights)

	for _, h := range heights {
		fmt.Printf("%d ", h)
	}
	fmt.Println()
}
