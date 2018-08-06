package main

import (
	"encoding/json"
	"fmt"
)

func stringifyAndPrint(obj interface{}) {
	result, _ := json.Marshal(obj)
	fmt.Println(string(result))
	fmt.Println()
}

func judgePoint24(nums []int) bool {
	type Operator byte
	type Operand *float32
	const (
		Plus Operator = iota
		Multiply
		Minus
		Divide
	)

	type TreeNode struct {
		Left     *TreeNode
		Right    *TreeNode
		Operator *Operator
		Value    Operand
		ID       int
	}

	const targetSum = 24.0
	const opCount = 3

	var floatNums []float32

	for i := 0; i < len(nums); i++ {
		floatNums = append(floatNums, float32(nums[i]))
	}

	nextPermutation := func(array []int) bool {
		// Find longest non-increasing suffix
		i := len(array) - 1
		for i > 0 && array[i-1] >= array[i] {
			i--
		}
		// Now i is the head index of the suffix

		// Are we at the last permutation already?
		result := i > 0

		if result {
			// Let array[i - 1] be the pivot
			// Find rightmost element that exceeds the pivot
			j := len(array) - 1
			for array[j] <= array[i-1] {
				j--
			}
			// Now the value array[j] will become the new pivot
			// Assertion: j >= i

			// Swap the pivot with j
			array[i-1], array[j] = array[j], array[i-1]

			// Reverse the suffix
			j = len(array) - 1
			for i < j {
				array[i], array[j] = array[j], array[i]
				i++
				j--
			}
		}

		return result
	}

	operatorNodes, valueNodes := (func(floatNums []float32) ([]TreeNode, []TreeNode) {
		operatorNodes := make([]TreeNode, 1)
		valueNodes := []TreeNode{}
		root := &(operatorNodes[0])

		for i := 0; i < len(nums); i++ {
			newNode := &TreeNode{Value: &floatNums[i]}

			if root.Left == nil {
				root.Left = newNode
			} else {
				if root.Right == nil {
					root.Right = newNode
				} else {
					rightNode := &TreeNode{Right: root, Left: newNode}
					operatorNodes = append(operatorNodes, *rightNode)
					root = rightNode
				}
			}

			valueNodes = append(valueNodes, *newNode)
		}
		return operatorNodes, valueNodes
	})(floatNums)

	stringifyAndPrint(len(valueNodes))
	stringifyAndPrint(len(operatorNodes))

	numIndices, operators := []int{}, make([]Operator, len(operatorNodes)+1)

	for i := 0; i < len(nums); i++ {
		numIndices = append(numIndices, i)
	}

	for ok := true; ok; ok = nextPermutation(numIndices) {
		for operators[0] == 0 {
			for i := 0; i < len(operatorNodes); i++ {
				operatorNodes[i].Operator = &operators[i+1]
			}

			for i := 1; i < len(numIndices); i++ {
				valueNodes[i].Value = &floatNums[numIndices[i]]
			}

			stringifyAndPrint(operatorNodes[0])

			i := len(operators) - 1
			operators[i]++
			for operators[i] > opCount && i > 0 {
				operators[i] = 0
				i--
				operators[i]++
			}
		}

		operators[0] = 0
	}

	return len(operatorNodes) > 0
}

func main() {
	//fmt.Println(judgePoint24([]int{0, 3, 0, 72}))
	fmt.Println(judgePoint24([]int{1, 3, 4, 6})) //6/(1-3/4)
	//fmt.Println(judgePoint24([]int{8, 3, 1, 2}))
	//fmt.Println(judgePoint24([]int{50, 2, 0, 1}))
	//fmt.Println(judgePoint24([]int{-50, 2, 0, -1}))
}
