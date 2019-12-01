package main

import "fmt"

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	res := []int{}

	jMax := tomatoSlices >> 2
	for j := 0; len(res) == 0 && j <= jMax; j++ {
		s := tomatoSlices - 4*j
		if s%2 > 0 {
			continue
		}
		s = s >> 1
		if cheeseSlices-j == s {
			res = append(res, j, s)
		}
	}

	return res
}

func main() {
	fmt.Println(numOfBurgers(16, 7))
	fmt.Println(numOfBurgers(17, 4))
	fmt.Println(numOfBurgers(0, 0))
	fmt.Println(numOfBurgers(2, 1))
	fmt.Println(numOfBurgers(4, 1))
}
