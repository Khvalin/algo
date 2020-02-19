package main

import (
	"bufio"
	"fmt"
  "sort"
	"os"
)

func toInt(buf []byte) uint32 {
	n := uint32(0)
	for _, v := range buf {
		n = n*10 + uint32(v-'0')
	}

	return n
}

func readData() (n, m, k uint32, connected [][]bool, special []uint32) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n = toInt(scanner.Bytes())

	scanner.Scan()
	m = toInt(scanner.Bytes())

	scanner.Scan()
	k = toInt(scanner.Bytes())

	special = make([]uint32, k)
	for i := range special {
		scanner.Scan()
		special[i] = toInt(scanner.Bytes())
	}
  _ = sort.Ints

	connected = make([][]bool, 1+n)
	for i := uint32(0); i < m; i++ {
		scanner.Scan()
		x := toInt(scanner.Bytes())

		scanner.Scan()
		y := toInt(scanner.Bytes())

		if connected[x] == nil {
			connected[x] = make([]bool, 1+n)
		}

		if connected[y] == nil {
			connected[y] = make([]bool, 1+n)
		}

		connected[x][y], connected[y][x] = true, true
	}

	return
}

func solve (n, m, k uint32, connected [][]bool, special []uint32) uint32 {
  return 0
}

func main() {
	fmt.Println(solve(readData()))
}
