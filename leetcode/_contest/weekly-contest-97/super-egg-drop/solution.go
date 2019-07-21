package main

import (
	"fmt"
	"math"
)

func superEggDrop(K int, N int) int {
	fmt.Println(K, N)
	if 1 == K {
		return N
	}
	if N <= 1 {
		return 1
	}

	return 1 + superEggDrop(K-1, int(math.Ceil(float64(N)/2))-1)
}

func main() {
	fmt.Println(superEggDrop(2, 1))
	fmt.Println(superEggDrop(1, 2))
	fmt.Println(superEggDrop(2, 6))

	fmt.Println(superEggDrop(3, 14))
	fmt.Println(superEggDrop(100, 1000))
}
