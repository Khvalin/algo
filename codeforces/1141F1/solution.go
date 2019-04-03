package main

import (
	"bufio"
	"fmt"
	"os"
)

type Range struct {
	start int
	end   int
}

func solve(a []int) []Range {
	L := len(a)

	sums := make([]map[int]Range, L)
	all := map[int]int{}

	for i := L - 1; i >= 0; i-- {
		cur := a[i]

		for j := i; j < L; j++ {
			if i != j {
				cur += a[j]
			}

			_, found := sums[i][cur]
			if !found {
				if sums[i] == nil {
					sums[i] = map[int]Range{}
				}

				all[cur]++

				sums[i][cur] = Range{start: i, end: j}
			}
		}
	}

	max := -1
	maxPath := []Range{}

	for len(all) > 0 {
		for cur := range all {
			i := 0
			count := 0
			path := []Range{}

			for i < L {
				sum, found := sums[i][cur]

				if !found {
					i++
					continue
				}
				delete(sums[i], cur)
				all[cur]--
				if all[cur] <= 0 {
					delete(all, cur)
				}

				count++
				path = append(path, sum)

				i = sum.end + 1

				if count > max {
					max = count
					maxPath = make([]Range, len(path))
					copy(maxPath, path)
				}
			}

			if cur == -5 {
				fmt.Println(path)
			}
		}
	}

	return maxPath
}

func toInt(buf []byte) int {
	n, mul := int(0), int(1)
	for i, v := range buf {
		if i == 0 && v == '-' {
			mul = int(-1)
			continue
		}
		n = n*10 + int(v-'0')
	}

	return mul * n
}

func main() {
	var n int

	fmt.Scan(&n)

	a := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := range a {
		scanner.Scan()
		a[i] = toInt(scanner.Bytes())
	}

	ans := solve(a)
	fmt.Println(len(ans))

	for _, pair := range ans {
		fmt.Printf("%v %v\n", 1+pair.start, 1+pair.end)
	}
}
