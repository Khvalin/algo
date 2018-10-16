package main

import (
	"fmt"
	"math"
)

func sumOfTwoPowers(n int, k int) [][]int {
	pows := map[int]bool{0: true}
	arr := []int{0}

	kF := float64(k)

	for i := 1; true; i++ {
		pow := int(math.Pow(float64(i), kF))
		if pow <= n {
			pows[pow] = true
			arr = append(arr, pow)
		} else {
			break
		}
	}

	//fmt.Println(arr)

	res := [][]int{}
	for _, pow := range arr {
		ok := pows[pow]

		if ok {
			b := n - pow
			found := pows[b]
			if found {
				res = append(res, []int{pow, b})
				delete(pows, pow)
				delete(pows, b)
			}
		}
	}

	return res
}

func main() {
	fmt.Println(sumOfTwoPowers(87539319, 3))
}
