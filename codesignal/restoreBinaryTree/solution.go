package main

import (
	"encoding/json"
	"fmt"
)

// Binary trees are already defined with this interface:
type Tree struct {
	Value interface{} `json:"value"`
	Left  *Tree `json:"left"`
	Right *Tree `json:"right"`
}

func restoreBinaryTree(inorder []int, preorder []int) *Tree {
	m := map[int]*Tree{}
	inPos := map[int]int{}

	for i, n := range inorder {
		inPos[n] = i
		m[n] = &Tree{Value: n}
	}

	lq := []int{0}
	rq := []int{0}

	next := 1
	for next < len(preorder) && (len(lq) > 0 || len(rq) > 0) {
		ind := -1
		pushToTop := len(lq) > 0

		if len(lq) > 0 {
			ind = lq[0]
			lq = lq[1:]
		} else {
			ind = rq[0]
			rq = rq[1:]
		}

		n := preorder[ind]
		node := m[n]

		push := func(next int) {
			if pushToTop {
				lq = append([]int{next}, lq...)
				rq = append(rq, next)
				return
			}

			lq = append(lq, next)
			rq = append([]int{next}, rq...)
		}

		if inPos[preorder[next]] < inPos[n]  && node.Left == nil {
			node.Left = m[preorder[next]]
			push(next)

			next++
		}

		if pushToTop {
			continue
		}

		if next < len(preorder) && inPos[preorder[next]] > inPos[n] && node.Right == nil {
			node.Right = m[preorder[next]]
			push(next)

			next++
		}
	}

	return m[preorder[0]]
}

func main() {
	inorder := []int{4, 2, 1, 5, 3, 6}
	preorder := []int{1, 2, 4, 3, 5, 6}

	inorder = []int{4, 2, 7, 1, 5, 3, 6}
	preorder = []int{1, 2, 4, 7, 3, 5, 6}

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
