package main

import "fmt"

func findDuplicate(nums []int) int {
	n := len(nums) - 1
	res := 1
	s := n * (n + 1) / 2
	total := 0
	for _, n := range nums {
		total += n
	}
	fmt.Println(s, total)
	return res
}

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate([]int{1, 2, 3, 4, 5, 6, 7, 8, 1, 1, 1}))
}
