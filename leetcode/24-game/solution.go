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

type operator byte
type operand *float32

const (
	//Plus const
	Plus operator = iota
	//Multiply const
	Multiply
	//Minus const
	Minus
	//Divide const
	Divide
)

type treeNode struct {
	Left     *treeNode `json:",omitempty"`
	Right    *treeNode `json:",omitempty"`
	Operator operator
	Value    operand `json:",omitempty"`
	//	ID       int
}

func (node treeNode) calc() float32 {
	var result float32
	if node.Value == nil {
		result = (node.Left).calc()
		rightValue := (node.Right).calc()

		switch node.Operator {
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

	operatorNodes, valueNodes, root := (func(floatNums []float32) ([]*treeNode, []*treeNode, *treeNode) {
		valueNodes := []*treeNode{}
		root := &treeNode{}
		operatorNodes := []*treeNode{}

		for i := 0; i < len(nums); i++ {
			newNode := &treeNode{Value: &floatNums[i]}

			if root.Left == nil {
				root.Left = newNode
			} else {
				if root.Right == nil {
					root.Right = newNode
				} else {
					rightNode := &treeNode{Right: root, Left: newNode}
					operatorNodes = append(operatorNodes, rightNode)
					root = rightNode
				}
			}

			valueNodes = append(valueNodes, newNode)
		}
		operatorNodes = append(operatorNodes, root)
		return operatorNodes, valueNodes, root
	})(floatNums)

	// fmt.Printf("%p", root)
	// fmt.Println(operatorNodes)

	//	stringifyAndPrint((valueNodes))
	//	stringifyAndPrint((root))

	numIndices, operators := []int{}, []operator{}

	for i := 0; i < len(nums); i++ {
		numIndices = append(numIndices, i)
		operators = append(operators, 0)
	}

	for ok := true; ok; ok = nextPermutation(numIndices) {
		for operators[0] == 0 {
			cur := root
			for i := 0; i < len(operatorNodes); i++ {
				cur.Operator = (operators[i+1])
				cur = cur.Right
			}
			//	stringifyAndPrint((root))

			for i := 0; i < len(numIndices); i++ {
				valueNodes[i].Value = &floatNums[numIndices[i]]
				//	fmt.Printf("%f ", floatNums[numIndices[i]])
			}
			//fmt.Println()

			res := root.calc()
			if res == targetSum {
				//		stringifyAndPrint((root))
				return true
			}
			fmt.Println(res)
			//stringifyAndPrint(root)

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

	return false
}

func main() {
	fmt.Println(judgePoint24([]int{1, 3, 4, 6})) //6/(1-3/4)
	fmt.Println(judgePoint24([]int{1, 9, 1, 2}))
}
