package main

import (
	"encoding/json"
	"fmt"
)

// Tree -- Binary trees are already defined with this interface:
type Tree struct {
	Value interface{} `json:"value"`
	Left  *Tree       `json:"left"`
	Right *Tree       `json:"right"`
}

func restoreBinaryTree(inorder, preorder []int) *Tree {
	L := len(inorder)
	m := map[int]int{}
	for i, n := range inorder {
		m[n] = i
	}

	var split func(i, j, k int) (*Tree, int)
	split = func(i, j, k int) (*Tree, int) {
		if k >= L || i >= j {
			return nil, k - 1
		}

		cur := preorder[k]
		node := &Tree{Value: cur}
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

func main() {
	inorder := []int{4, 2, 1, 5, 3, 6}
	preorder := []int{1, 2, 4, 3, 5, 6}

	// inorder = []int{4, 2, 7, 1, 5, 3, 6}
	// preorder = []int{1, 2, 4, 7, 3, 5, 6}

	tree := restoreBinaryTree(inorder, preorder)

	r, _ := json.MarshalIndent(tree, "", "  ")
	fmt.Println(string(r))

	// inorder = []int{100}
	// preorder = []int{100}
	// //fmt.Printf("%v \n", restoreBinaryTree(inorder, preorder))

	// inorder = []int{4, 2, 1, 5, 3, 6}
	// preorder = []int{1, 2, 4, 3, 5, 6}
	// //	fmt.Printf("%v \n", restoreBinaryTree(inorder, preorder))

	// inorder = []int{-100000, -99999, -99998}
	// preorder = []int{-99999, -100000, -99998}
	// fmt.Printf("%v \n", restoreBinaryTree(inorder, preorder))
}
