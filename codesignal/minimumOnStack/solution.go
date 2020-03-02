package main

import (
	"fmt"
	"strconv"
	"strings"
)

func minimumOnStack(operations []string) []int {
	L := 1 + len(operations)
	res := make([]int, 0, L)
	min := make([]int, 0, L)
	stack := make([]int, 0, L)

	for _, op := range operations {
		if strings.HasPrefix(op, "push") {
			s := strings.TrimPrefix(op, "push ")
			d, _ := strconv.Atoi(s)
			stack = append(stack, d)
			if len(min) == 0 || d <= min[len(min)-1] {
				min = append(min, d)
			}
		}

		if op == "min" {
			res = append(res, min[len(min)-1])
		}

		if strings.HasPrefix(op, "pop") {
			d := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if min[len(min)-1] == d {
				min = min[:len(min)-1]
			}
		}
	}

	return res
}

func main() {
	fmt.Println(minimumOnStack([]string{"push 10",
		"min",
		"push 5",
		"min",
		"push 8",
		"min",
		"pop",
		"min",
		"pop",
		"min",
	}))
}
