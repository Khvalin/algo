package main

import (
	"fmt"
)

func judgePoint24(nums []int) bool {
	const targetSum = 24.0
	const opCount = 3
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
			temp := array[i-1]
			array[i-1] = array[j]
			array[j] = temp

			// Reverse the suffix
			j = len(array) - 1
			for i < j {
				temp := array[i]
				array[i] = array[j]
				array[j] = temp
				i++
				j--
			}
		}

		return result
	}

	calc := func(numIndices []int, opIndices []int) float64 {
		result := float64(nums[numIndices[0]])
		for i := 1; i < len(numIndices); i++ {
			cur := float64(nums[numIndices[i]])
			switch opIndices[i] {
			case 0:
				result += cur
			case 1:
				result -= cur

			case 2:
				result *= cur

			case 3:
				result /= cur
			}
		}

		return result
	}

	numIndices, opIndices := []int{}, []int{}

	for i := 0; i < len(nums); i++ {
		numIndices = append(numIndices, i)
		opIndices = append(opIndices, 0)
	}

	for ok := true; ok; ok = nextPermutation(numIndices) {
		for opIndices[0] == 0 {
			calcResult := calc(numIndices, opIndices)
			fmt.Println(numIndices, opIndices, calcResult)
			if targetSum == calcResult {

				return true
			}

			i := len(opIndices) - 1
			opIndices[i]++
			for opIndices[i] > opCount && i > 0 {
				opIndices[i] = 0
				i--
				opIndices[i]++
			}
		}

		opIndices[0] = 0
	}

	return false
}

func main() {
	//fmt.Println(judgePoint24([]int{0, 3, 0, 72}))
	fmt.Println(judgePoint24([]int{1, 3, 4, 6})) //6/(1-3/4)
	//fmt.Println(judgePoint24([]int{8, 3, 1, 2}))
	//fmt.Println(judgePoint24([]int{50, 2, 0, 1}))
	//fmt.Println(judgePoint24([]int{-50, 2, 0, -1}))
}
