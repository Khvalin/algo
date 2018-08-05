package main

import (
	"encoding/json"
	"fmt"
)

func judgePoint24(nums []int) bool {
	type Operator byte
	type Operand *float32
	const (
		Any Operator = iota
		Plus
		Multiply
		Minus
		Divide
	)

	type TreeNode struct {
		Left     *TreeNode
		Right    *TreeNode
		Operator *Operator
		Operand  Operand
		ID       int
	}

	const targetSum = 24.0

	nodes := (func() []TreeNode {
		nodes := make([]TreeNode, 1)
		for i := 0; i < len(nums); i++ {
			floatValue := float32(nums[i])
			newNode := &TreeNode{Operand: &floatValue}

			operandNode := &(nodes[0])

			if operandNode.Left == nil {
				operandNode.Left = newNode
			} else {
				if operandNode.Right == nil {
					operandNode.Right = newNode
				}
			}

			nodes = append(nodes, *newNode)
		}
		return nodes
	})()

	json, _ := json.Marshal(nodes)
	output := string(json)
	fmt.Println(output)

	return false
}

func main() {
	//fmt.Println(judgePoint24([]int{0, 3, 0, 72}))
	fmt.Println(judgePoint24([]int{1, 3, 4, 6})) //6/(1-3/4)
	//fmt.Println(judgePoint24([]int{8, 3, 1, 2}))
	//fmt.Println(judgePoint24([]int{50, 2, 0, 1}))
	//fmt.Println(judgePoint24([]int{-50, 2, 0, -1}))
}
