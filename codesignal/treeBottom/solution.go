package main

import "fmt"

func treeBottom(tree string) []int {
	s := [2]int{1, 0}

	max := 0
	res := []int{}
	stack := [][2]int{s}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		i, level := node[0], node[1]

		fmt.Print(tree[i])

		switch {
		case tree[i] == ' ':
			{
				stack[len(stack)-1][0]++
			}
		case tree[i] >= '0' && tree[i] <= '9':
			{
				n := 0

				for ; tree[i] != ' '; i++ {
					n *= 10
					n += int(tree[i] - '0')
				}

				if level > max {
					max = level
					res = []int{}
				}

				if level == max {
					res = append(res, n)
				}

				stack[len(stack)-1][0] = i + 1
			}

		case tree[i] == '(':
			{
				stack = append(stack, [2]int{i + 1, level + 1})

			}

		case tree[i] == ')':
			{
				stack = stack[:len(stack)-1]
			}
		}
	}

	return res
}

func main() {
	r := treeBottom("(2 (7 (2 () ()) (6 (5 () ()) (11 () ()))) (5 () (9 (4 () ()) ())))")
	fmt.Println(r)
}
