package main

import (
	"encoding/json"
	"fmt"
)

// Binary trees are already defined with this interface:
type Tree struct {
	Value interface{}
	Left  *Tree
	Right *Tree
}

func restoreBinaryTree(inorder []int, preorder []int) *Tree {
	m := map[int]bool{}
	var root, subroot *Tree
	var last (*Tree)
	i := 0
	leap := false

	for _, n := range inorder {
		node := &Tree{Value: n}
		if !leap {
			node.Left = last
			leap = false
		}

		if i < len(preorder) && n == preorder[i] {

			if i == 0 {
				root = node
			}
			i++
			for i < len(preorder) && (m[preorder[i]]) {
				i++
			}

			if subroot != nil {
				subroot.Right = node
			} else {
				leap = true
			}

			subroot = node
		}

		last = node
		m[n] = true
	}

	return root
}

func main() {
	inorder := []int{4, 2, 1, 5, 3, 6}
	preorder := []int{1, 2, 4, 3, 5, 6}
	tree := restoreBinaryTree(inorder, preorder)

	r, _ := json.Marshal(tree)
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
