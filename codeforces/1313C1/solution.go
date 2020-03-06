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

//func calc()

func solve(heights []uint64) []uint64 {
	L := len(heights)
	res, t := make([]uint64, L), make([]uint64, L)
	sums := make([][2]map[uint64]uint64, L)
	for i := range sums {
		sums[i][0], sums[i][1] = map[uint64]uint64{}, map[uint64]uint64{}
	}

	var max uint64

	for i, h := range heights {
		copy(t, heights)

		var s uint64
		var l, r uint64

    var _, store = sums[i][0][h]
    store = !store
		for j := i - 1; j >= 0; j-- {
			if t[j] > t[j+1] {
				t[j] = t[j+1]
			}

			if s, f := sums[j][0][t[j]]; f {
				l += s
				//store = false
				break
			}

			l += t[j]
		}

		if store {
			for j := 1; j < i; j++ {
				sums[j][0][t[j]] =  t[j-1]
				if j > 1 {
					sums[j][0][t[j]] += sums[j][0][t[j-2]]
				}
			}
		}

		_, store = sums[i][1][h]
    store = !store
		for j := i + 1; j < L; j++ {
			if t[j] > t[j-1] {
				t[j] = t[j-1]
			}

			if s, f := sums[j][1][t[j]]; f {
				r += s
				//store = false
				break
			}

			r += t[j]
		}

		if store {
			for j := L - 2; j > i; j-- {
				sums[j][1][t[j]] = t[j+1]
				if j < L-2 {
					sums[j][1][t[j]] += sums[j][1][t[j+2]]
				}
			}
		}

		sums[i][0][h] = l
		sums[i][1][h] = r

		s = l + r + h

		if s > max {
			max = s
		}
	}

	fmt.Println(max)

	return res
}

func writeData(heights []uint64) {
	stdout := bufio.NewWriterSize(os.Stdout, 4096)
	defer stdout.Flush()

	for _, h := range heights {
		fmt.Fprintf(stdout, "%d ", h)
	}
	fmt.Fprintf(stdout, "\n")
}

func main() {
	heights := readData()
	heights = solve(heights)
//	writeData(heights)
}
