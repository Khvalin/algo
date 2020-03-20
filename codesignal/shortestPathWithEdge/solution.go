package main

import "fmt"

func shortestPathWithEdge(start, finish, weight int, graph [][]int) int {
	const nmax = 9223372036854775807

	type state struct {
		ind int
		sum int
		fix int
	}

	minLen := [2][500]int{}

	for i := range minLen {
		for j := range minLen[i] {
			minLen[i][j] = nmax
		}
	}

	q := []state{state{ind: start - 1, sum: 0, fix: 0}}
	r := nmax

	for len(q) > 0 {
		st := q[0]
		q = q[1:]

		if st.ind == finish-1 {
			if st.sum < r {
				r = st.sum
			}
			continue
		}

		if minLen[st.fix][st.ind] < st.sum {
			continue
		}

		minLen[st.fix][st.ind] = st.sum

		for i, w := range graph[st.ind] {
			if w > 0 && minLen[st.fix][i] > w+st.sum {
				q = append(q, state{ind: i, sum: w + st.sum, fix: st.fix})
			}

			if w == 0 && st.fix == 0 && minLen[st.fix][i] > weight+st.sum {
				q = append(q, state{ind: i, sum: weight + st.sum, fix: 1})
			}
		}
	}

	return r
}

func main() {
	start, finish, weight := 1, 4, 2
	graph := [][]int{
		{0, 2, 0, 4, 0},
		{2, 0, 1, 0, 0},
		{0, 1, 0, 3, 0},
		{4, 0, 3, 0, 1},
		{0, 0, 0, 1, 0},
	}

	fmt.Println(shortestPathWithEdge(start, finish, weight, graph))
}
