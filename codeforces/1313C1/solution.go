package main

import (
	"bufio"
	"fmt"
	"os"
)

func toInt(buf []byte) int {
	n := int(0)
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}

	return n
}

func readData() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n := toInt(scanner.Bytes())

	r := make([]int, n)
	for i := range r {
		scanner.Scan()
		r[i] = toInt(scanner.Bytes())
	}

	return r
}

func solve(heights []int) []int {
	L := len(heights)

	mins := make([]int, L)
	mins[0] = heights[0]
	for i := 1; i < L; i++ {
		mins[i] = mins[i-1]
		if /*(i >= L-1 || heights[i] < heights[i+1]) && */ i <= 0 || heights[i] < heights[i-1] {
			mins[i] = heights[i]
		}
	}

	fmt.Println(mins)

	type state struct {
		ind int
		max int
	}

	type result struct {
		sum int
		ans []int
	}

	memo := map[state]result{}
	_ = memo
	start := state{-1, heights[0]}
	_ = start

	var calc func(i, max int) result
	calc = func(i, max int) result {
		if i < 0 {
			return result{0, []int{}}
		}

		// r, f := memo[st]
		// if f {
		// 	return r
		// }
		r := result{}

		return r
	}

	r := calc(L-1, heights[L-1])

	return r.ans
}

func main() {
	heights := readData()
	heights = solve(heights)

	for _, h := range heights {
		fmt.Printf("%v ", h)
	}
	fmt.Println()
}
