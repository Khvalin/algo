package main

import "fmt"

func deleteTreeNodes(nodes int, parent []int, value []int) int {
	sum := make([]int, nodes)
	copy(sum, value)

	for i := nodes - 1; i >= 0; i-- {
		p := parent[i]
		if p >= 0 {
			sum[p] += sum[i]
		}
	}

	res := 0
	for i := range sum {
		p := parent[i]

		if p >= 0 && sum[p] == 0 {
			sum[i] = 0
		}

		if sum[i] != 0 {
			res++
		}
	}

	return res
}

func main() {
	nodes, parent, value := 7, []int{-1, 0, 0, 1, 2, 2, 2}, []int{1, -2, 4, 0, -2, -1, -1}
	fmt.Println(deleteTreeNodes(nodes, parent, value))
}
