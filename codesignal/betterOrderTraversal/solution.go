package main

import (
	"sort"
)

// Definition for binary tree:
type Tree struct {
	Value interface{}
	Left  *Tree
	Right *Tree
}

func betterOrderTraversal(root *Tree) []int {
	var inOrder func(root *Tree) []int
	inOrder = func(root *Tree) []int {
		res := []int{}
		if root != nil {
			res = inOrder(root.Left)
			res = append(res, root.Value.(int))
			res = append(res, inOrder(root.Right)...)
		}
		return res
	}

	var preOrder func(root *Tree) []int
	preOrder = func(root *Tree) []int {
		res := []int{}
		if root != nil {
			res = append(res, root.Value.(int))
			res = append(res, preOrder(root.Left)...)
			res = append(res, preOrder(root.Right)...)
		}
		return res
	}

	var postOrder func(root *Tree) []int
	postOrder = func(root *Tree) []int {
		res := []int{}
		if root != nil {
			res = append(res, postOrder(root.Left)...)
			res = append(res, postOrder(root.Right)...)
			res = append(res, root.Value.(int))
		}
		return res
	}

	orders := [][]int{inOrder(root), preOrder(root), postOrder(root)}

	sort.Slice(orders, func(i, j int) bool {
		n := len(orders[i])
		if n > len(orders[j]) {
			n = len(orders[j])
		}
		k := 0
		for k < n-1 && (orders[i][k] == orders[j][k]) {
			k++
		}

		return orders[i][k] < orders[j][k]
	})

	// fmt.Println(orders)

	return orders[0]
}
