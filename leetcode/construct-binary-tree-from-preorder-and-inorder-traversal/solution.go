

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	L := len(inorder)
	m := map[int]int{}
	for i, n := range inorder {
		m[n] = i
	}

	var split func(i, j, k int) (*TreeNode, int)
	split = func(i, j, k int) (*TreeNode, int) {
		if k >= L || i >= j {
			return nil, k - 1
		}

		cur := preorder[k]
		node := &TreeNode{Val: cur}
		if j == i+1 {
			return node, k
		}

		if k < L-1 && m[cur] > m[preorder[k+1]] {
			node.Left, k = split(i, m[cur], k+1)
		}
		if k < L-1 && m[cur] < m[preorder[k+1]] {
			node.Right, k = split(m[cur]+1, j, k+1)
		}

		return node, k
	}

	r, _ := split(0, L, 0)

	return r
}

/*
	[1, 2, 4, 7, 3, 5, 6]
[4, 2, 7, 1, 5, 3, 6]
*/