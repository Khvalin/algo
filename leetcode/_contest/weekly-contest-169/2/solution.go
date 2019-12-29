package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var inOrder func(node *TreeNode, res *[]int)
	inOrder = func(node *TreeNode, res *[]int) {
		if node == nil {
			return
		}

		inOrder(node.Left, res)
		*res = append(*res, node.Val)
		inOrder(node.Right, res)
	}

	a1 := []int{}
	a2 := []int{}

	inOrder(root1, &a1)
	inOrder(root2, &a2)

	res := make([]int, len(a2)+len(a1))

	i, j, k := 0, 0, 0
	for i < len(a1) || j < len(a2) {
		if i >= len(a1) || (j < len(a2) && a2[j] <= a1[i]) {
			res[k] = a2[j]
			k++
			j++
			continue
		}

		res[k] = a1[i]
		k++
		i++
	}

	return res
}
