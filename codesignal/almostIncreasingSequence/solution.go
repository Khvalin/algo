package main

import (
	"fmt"
	"math"
)

func almostIncreasingSequence(sequence []int) bool {
	sequence = append([]int{-int(math.Pow10(5) + 1)}, sequence...)
	sequence = append(sequence, int(math.Pow10(5))+1)

	start := 1
	end := len(sequence) - 1

	for start < end && sequence[start] > sequence[start-1] {
		start++
	}

	for end > start && sequence[end] > sequence[end-1] {
		end--
	}

	if end-start <= 2 {
		//count := 0
		for i := start - 1; i <= end; i++ {
			s := i - 2
			if s < 1 {
				s = 1
			}
			e := i + 2
			if e > len(sequence) {
				e = len(sequence)
			}

			fmt.Println(s, e)

			result := true
			for j := s + 1; j <= e; j++ {
				left, right := j-1, j
				if left == i {
					left--
				}
				if right == i {
					right++
				}
				result = result && (left < 0 || right >= len(sequence) || (sequence[right] > sequence[left]))
			}

			if result {
				return true
			}

		}

	}

	return false
}

func main() {
	fmt.Println(almostIncreasingSequence([]int{1, 3, 2, 1}))
	fmt.Println(almostIncreasingSequence([]int{1, 2, 1, 2}))
	fmt.Println(almostIncreasingSequence([]int{1, 2, 5, 3, 5}))
	fmt.Println(almostIncreasingSequence([]int{1, 2, 3, 4, 3, 6}))
	fmt.Println(almostIncreasingSequence([]int{1, 2, 5, 5, 5}))
}
