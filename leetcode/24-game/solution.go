package main

import (
	"encoding/json"
	"fmt"
	"math"
)

func stringifyAndPrint(obj interface{}) {
	result, _ := json.Marshal(obj)
	fmt.Println(string(result))
	fmt.Println()
}

type operator byte
type operand *float32

const (
	//Plus const
	Plus operator = iota
	//Multiply const
	Multiply = 1
	//Minus const
	Minus = 2
	//Divide const
	Divide = 3
)

type treeNode struct {
	Left     *treeNode `json:",omitempty"`
	Right    *treeNode `json:",omitempty"`
	Operator *operator
	Value    operand `json:",omitempty"`
}

func (node *treeNode) traverse(visit func(*treeNode)) {
	if node != nil {
		visit(node)

		if (node).Right != nil {
			node.Right.traverse(visit)
		}

		if (node).Left != nil {
			node.Left.traverse(visit)
		}
	}
}

func (node *treeNode) calc() float32 {
	var result float32
	if node.Value == nil {
		result = (node.Left).calc()
		rightValue := (node.Right).calc()

		switch *(node.Operator) {
		case Plus:
			result += rightValue
		case Minus:
			result -= rightValue
		case Multiply:
			result *= rightValue
		case Divide:
			result /= rightValue
		}

	} else {
		result = *(node.Value)
	}

	return result
}

func judgePoint24(nums []int) bool {

	const targetSum = float64(24)
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

	createLopsidedTree := (func(floatNums []float32) *treeNode {
		root := &treeNode{}

		for i := 0; i < len(nums); i++ {
			newNode := &treeNode{Value: &floatNums[i]}

			if root.Left == nil {
				root.Left = newNode
			} else {
				if root.Right == nil {
					root.Right = newNode
				} else {
					root = &treeNode{Right: root, Left: newNode}
				}
			}
		}

		return root
	})

	createBalancedTree := (func(floatNums []float32) *treeNode {
		nodes := [](*treeNode){}

		for i := 0; i < len(nums); i++ {
			nodes = append(nodes, &treeNode{Value: &floatNums[i]})
		}

		for len(nodes) > 1 {
			left, right := nodes[0], nodes[1]
			nodes = nodes[2:]

			parent := &treeNode{Left: left, Right: right}
			nodes = append(nodes, parent)
		}

		return nodes[0]
	})

	roots := []*treeNode{createLopsidedTree(floatNums), createBalancedTree(floatNums)}

	numIndices, operators := []int{}, []operator{}

	for i := 0; i < len(nums); i++ {
		numIndices = append(numIndices, i)
		operators = append(operators, 0)
	}

	for ok := true; ok; ok = nextPermutation(numIndices) {
		operators[0] = 0
		for operators[0] == 0 {
			//	fmt.Println(operators)
			for k := 0; k < len(roots); k++ {
				root := roots[k]

				i, j := 0, 1
				root.traverse(func(node *treeNode) {
					if node.Right == nil && node.Left == nil {
						node.Value = &floatNums[numIndices[i]]
						i++
					} else {
						node.Operator = &(operators[j])
						j++
					}
				})

				res := float64(root.calc())
				if math.Abs(targetSum-res) < 0.001 {
					return true
				}

			}

			i := len(operators) - 1
			operators[i]++
			for operators[i] > opCount && i > 0 {
				operators[i] = 0
				i--
				operators[i]++
			}
		}
	}

	return false
}

func main() {
	fmt.Println(judgePoint24([]int{1, 3, 4, 6})) //6/(1-3/4)
	fmt.Println(judgePoint24([]int{1, 9, 1, 2}))
	fmt.Println(judgePoint24([]int{4, 1, 8, 7}))
	fmt.Println(judgePoint24([]int{1, 2, 1, 2}))

	fmt.Println(judgePoint24([]int{3, 3, 8, 8})) // (8*3/3)/8

}
