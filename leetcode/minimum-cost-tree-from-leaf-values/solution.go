package main

import "fmt"

// The value of each non-leaf node is equal to the product of the largest leaf value in its left and right subtree respectively.
func mctFromLeafValues(arr []int) int {
	const inf = 4001 * 19
	const nmax = 42

	type tree struct {
		max, sum int
	}
  
  zeroTree := tree{}

	combineTrees := func(left, right tree) tree {
		left.sum = right.max*left.max + left.sum + right.sum

		if right.max > left.max {
			left.max = right.max
		}

		return left
	}

  l := len(arr)
	memo := make([][]tree, l + 1)
  memo[0] = make([]tree, l + 1)
	for i := 1; i < l; i++ {
    memo[i] = make([]tree, l + 1)


		val := arr[i-1] * arr[i]

		t := tree{max: arr[i], sum: val}
		if t.max < arr[i-1] {
			t.max = arr[i-1]
		}
		memo[i-1][i] = t

		memo[i][i] = tree{max: arr[i]}
	}
  
  memo[0][0] = tree{max: arr[0]}

	infTree := tree{max: inf, sum: inf}

	var solve func(left, right int) tree
	solve = func(left, right int) tree {
    res := memo[left][right]
		if res != zeroTree {
			return res
		}

		res = infTree
		for i := left; i <= right; i++ {
			m := memo[i][i]

			var rSub, lSub tree

			if i > left {
				lSub = solve(left, i-1)
			}

			if i < right {
				rSub = solve(i+1, right)
			}

			if v := combineTrees(combineTrees(lSub, m), rSub); v.sum < res.sum {
				res = v
			}

			if v := combineTrees(lSub, combineTrees(m, rSub)); v.sum < res.sum {
				res = v
			}

		}

		memo[left][right] = res

		return res
	}

	tr := solve(0, len(arr)-1)

	return tr.sum
}

func main() {
	fmt.Println(32, mctFromLeafValues([]int{6, 2, 4}))
	fmt.Println(15, mctFromLeafValues([]int{1, 3, 3, 1}))
	fmt.Println(541, mctFromLeafValues([]int{15, 10, 2, 3, 4, 1, 1, 1, 5, 4, 3, 6, 15}))
	fmt.Println(5, mctFromLeafValues([]int{2, 1, 1, 1, 1}))

	fmt.Println(19, mctFromLeafValues([]int{3, 1, 1, 1, 2, 3}))
}
