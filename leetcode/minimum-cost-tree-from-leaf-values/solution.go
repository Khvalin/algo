package main

import "fmt"

// The value of each non-leaf node is equal to the product of the largest leaf value in its left and right subtree respectively.
func mctFromLeafValues(arr []int) int {
	const inf = 4001 * 19

	type state struct {
		left, right int
	}

	type tree struct {
		max, sum int
	}

	zeroTree := tree{max: -1, sum: -1}

	combineTrees := func(left, right tree) tree {
		if left == zeroTree {
			return right
		}

		if right == zeroTree {
			return left
		}

		max := left.max
		sum := right.max*left.max + left.sum + right.sum
		if sum > inf {
			sum = inf
		}

		if right.max > max {
			max = right.max
		}

		return tree{max: max, sum: sum}
	}

	memo := map[state]tree{}
	for i := 1; i < len(arr); i++ {
		val := arr[i-1] * arr[i]

		t := tree{max: arr[i], sum: val}
		if t.max < arr[i-1] {
			t.max = arr[i-1]
		}
		memo[state{i - 1, i}] = t

		memo[state{i, i}] = tree{max: arr[i]}
	}
	memo[state{0, 0}] = tree{max: arr[0]}

	infTree := tree{max: inf, sum: inf}

	var solve func(left, right int) tree
	solve = func(left, right int) tree {
		if left < 0 || right >= len(arr) {
			return zeroTree
		}

		st := state{left, right}

		if res, found := memo[st]; found {
			return res
		}

		res := infTree
		for i := left; i < right; i++ {
			mid := solve(i, i+1)

			leftSibling, rightSibling := zeroTree, zeroTree
			leftSubtree, rightSubtree := zeroTree, zeroTree

			if i > left {
				leftSibling = solve(i-1, i-1)
				if i > left+1 {
					leftSubtree = solve(left, i-2)
				}
			}

			if i < right-1 {
				rightSibling = solve(i+2, i+2)
				if i < right-2 {
					rightSubtree = solve(i+2, right)
				}
			}

			variants := []tree{
				combineTrees(combineTrees(combineTrees(leftSubtree, leftSibling), mid), combineTrees(rightSibling, rightSubtree)),
				combineTrees(combineTrees(leftSubtree, combineTrees(leftSibling, mid)), combineTrees(rightSibling, rightSubtree)),
				combineTrees(combineTrees(leftSubtree, leftSibling), combineTrees(combineTrees(mid, rightSibling), rightSubtree)),
				combineTrees(combineTrees(leftSubtree, leftSibling), combineTrees(mid, combineTrees(rightSibling, rightSubtree))),
			}

			for _, v := range variants {
				if v.sum < res.sum {
					res = v
				}
			}
		}

		memo[st] = res

		return res
	}

	tr := solve(0, len(arr)-1)

	return tr.sum
}

func main() {
	fmt.Println(32, mctFromLeafValues([]int{6, 2, 4}))
	fmt.Println(15, mctFromLeafValues([]int{1, 3, 3, 1}))
	fmt.Println(541, mctFromLeafValues([]int{15, 10, 2, 3, 4, 1, 1, 1, 5, 4, 3, 6, 15}))
}
