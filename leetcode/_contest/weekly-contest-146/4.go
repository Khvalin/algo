package main

import "fmt"

func maxAbsValExpr(arr1 []int, arr2 []int) int {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}

		return a
	}

	r := -9999999999999999
	for i := range arr1 {
		for j := range arr2 {
			n := abs(arr1[i]-arr1[j]) + abs(arr2[i]-arr2[j]) + abs(i-j)

			if n > r {
				r = n
			}
		}
	}

	return r
}

func main() {
	fmt.Println(maxAbsValExpr([]int{1, -2, -5, 0, 10},
		[]int{0, -2, -1, -7, -4}))
}
