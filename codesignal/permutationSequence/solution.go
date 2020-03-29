package main

import "fmt"

func permutationSequence(n int, k int) string {
	a := []rune{}
	for n > 0 {
		a = append([]rune{rune('0' + n%10)}, a...)
		n /= 10
	}

	return string(a)
}

func main() {
	fmt.Println(permutationSequence(123, 3))
}
