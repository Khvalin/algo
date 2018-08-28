package main

// Tree Definition for binary tree:
type Tree struct {
	Value interface{}
	Left  *Tree
	Right *Tree
}

func mostFrequentSum(t *Tree) []int {
	//sums := []int{}

	var visit func(node *Tree) []int

	visit = func(node *Tree) []int {
		result := []int{}
		leftSums := []int{}
		rightSums := []int{}
		total := 0
		total, _ = (node.Value).(int)

		if node.Left != nil {
			leftSums := visit(node.Left)
			total += leftSums[len(leftSums)-1]
		}

		if node.Right != nil {
			rightSums := visit(node.Right)
			total += rightSums[len(rightSums)-1]
		}

		result = append(leftSums, rightSums...)

		return append(result, total)
	}

	return visit(t)
}

func main() {
	//
}
