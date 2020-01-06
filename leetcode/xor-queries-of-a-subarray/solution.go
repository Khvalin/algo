package main

import "fmt"

func xorQueries(arr []int, queries [][]int) []int {
	xors := make([]int, 1+len(arr))
	for i := 0; i < len(arr); i++ {
		xors[i+1] = xors[i] ^ arr[i]
	}

	res := make([]int, len(queries))
	for i, q := range queries {
		res[i] = xors[q[1]+1] ^ xors[q[0]]
	}

	return res
}

func main() {
	arr := []int{4, 8, 2, 10}
	queries := [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}
	res := xorQueries(arr, queries)
	fmt.Println(res)
}
