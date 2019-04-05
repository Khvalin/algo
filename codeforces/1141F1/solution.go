package main

import (
	"bufio"
	"fmt"
	"os"
)

//Range ...
type Range struct {
	start int
	end   int
}

func solve(a []int) []Range {
	L := len(a)

	sums := make([]map[int]int, L)
	//	all := map[int]bool{}

	for i := L - 1; i >= 0; i-- {
		cur := 0

		for j := i; j < L; j++ {
			cur += a[j]

			//			all[cur] = true

			_, found := sums[i][cur]
			if !found {
				if sums[i] == nil {
					sums[i] = map[int]int{}
				}

				sums[i][cur] = j
			}
		}
	}

	memo := map[[2]int][]int{}

	var calc func(start, sum int) []int
	calc = func(start, sum int) []int {
		hash := [2]int{start, sum}
		out, found := memo[hash]

		if found {
			return out
		}

		out = []int{}

		if start < L {
			if next, found := sums[start][sum]; found {
				max := []int{start}
				for i := 1 + next; i < L; i++ {

					if _, found = sums[i][sum]; found {
						p := calc(i, sum)

						if len(p) >= len(max) {
							max = append([]int{start}, p...)
						}
					}
				}

				out = make([]int, len(max))
				copy(out, max)
			}
		}

		memo[hash] = out

		return out
	}

	target := 0
	max := -1
	var starts []int
	for i := L - 1; i >= 0; i-- {
		for s := range sums[i] {
			c := calc(i, s)
			if len(c) > max {
				target = s
				max = len(c)

				starts = c
			}
		}
	}

	res := make([]Range, max)

	for i, n := range starts {
		res[i] = Range{n, sums[n][target]}
	}

	return res
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
