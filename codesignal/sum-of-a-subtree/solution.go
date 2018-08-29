package main

import (
	"fmt"
	"sort"
)

// Tree Definition for binary tree:
type Tree struct {
	Value interface{}
	Left  *Tree
	Right *Tree
}

func mostFrequentSum(t *Tree) []int {
	var visit func(node *Tree) []int

	visit = func(node *Tree) []int {
		if node == nil {
			return []int{}
		}

		result := []int{}
		leftSums := []int{}
		rightSums := []int{}
		total := 0
		total, _ = (node.Value).(int)

		if node.Left != nil {
			leftSums = visit(node.Left)
			total += leftSums[0]
		}

		if node.Right != nil {
			rightSums = visit(node.Right)
			total += rightSums[0]
		}

		result = append(result, total)
		result = append(result, leftSums...)
		result = append(result, rightSums...)

		//fmt.Println(result)
		return result
	}

	sums := visit(t)
	sort.Ints(sums)

	freq := make(map[int]int)
	maxFreq := 0

	for i := 0; i < len(sums); i++ {
		_, ok := freq[sums[i]]
		if !ok {
			freq[sums[i]] = 0
		}

		freq[sums[i]]++

		if freq[sums[i]] > maxFreq {
			maxFreq = freq[sums[i]]
		}
	}

	result := []int{}
	for i := 0; i < len(sums); i++ {
		if freq[sums[i]] == maxFreq && (i == 0 || sums[i-1] != sums[i]) {
			result = append(result, sums[i])
		}
	}

	return result
}

func main() {
	root := &Tree{Value: -2, Left: &Tree{Value: -3}, Right: &Tree{Value: 2}}
	/*
			t: {
		    "value": 5,
		    "left": {
		        "value": 2,
		        "left": null,
		        "right": null
		    },
		    "right": {
		        "value": -3,
		        "left": {
		            "value": 10,
		            "left": null,
		            "right": null
		        },
		        "right": {
		            "value": 4,
		            "left": null,
		            "right": null
		        }
		    }
		}
	*/

	root2 := &Tree{Value: 5, Left: &Tree{Value: 2}, Right: &Tree{Value: -3, Left: &Tree{Value: 10}, Right: &Tree{Value: 4}}}
	root3 := &Tree{Value: 42}

	fmt.Println(mostFrequentSum(root))
	fmt.Println(mostFrequentSum(root2))
	fmt.Println(mostFrequentSum(root3))
}
