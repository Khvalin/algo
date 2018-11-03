package main

import (
	"fmt"
)

//I = int
type I = int

func jumpingJimmy2(T []I, P []I, O []I, h I) (r I) {
	i, L, N := -1, map[I]I{}, len(T)
	t := N > 0

	for _, p := range P {
		L[p] = 1
	}

	for _, p := range O {
		L[p] = i
	}

	c := t

	for c {
		c = !t

		s, l := h, i+1
		for l < N && s >= T[l] {
			s -= T[l]

			if s >= 0 {
				i = l
				r += T[l]

				c = t
			}
			l++
		}

		h += L[i]
	}

	return
}

func main() {
	T := []int{1, 4, 3, 2, 2, 2, 2, 1, 4, 4, 2, 3, 3, 4, 1, 4, 2, 1, 2, 4, 1, 2, 2, 4, 1}
	P := []int{1, 3, 11}
	O := []int{2, 4, 5, 7, 12, 20, 22}

	fmt.Println(jumpingJimmy2(T, P, O, 4), 56)

	T = []int{2, 6, 2, 2, 2, 3, 8, 9, 7, 5, 2, 3, 5, 7, 3, 3, 2, 4, 4, 7, 5, 5, 5, 7, 4, 5, 10, 2, 10, 10, 3, 5, 2, 9, 2, 6, 7, 8, 2, 8, 8, 3, 6, 8, 9, 3, 2, 5, 6, 6, 4, 2, 1, 4, 1, 8, 9, 2, 4, 9, 9, 5, 2, 6, 3, 5, 5, 3, 3, 3, 7, 8, 9, 7, 5, 4, 10, 1, 4, 1, 2, 9, 2, 7, 3, 6, 1, 4, 9, 6, 8, 9, 7, 3, 5, 2, 8}
	P = []int{4, 10, 12, 19, 26, 33, 35, 40, 43, 46, 47, 51, 69, 71, 73, 74, 75, 79, 80, 93}
	O = []int{1, 3, 11, 50, 56, 65, 81, 88, 94}
	fmt.Println(jumpingJimmy2(T, P, O, 10), 492)

	T = []int{}
	P = []int{}
	O = []int{}
	fmt.Println(jumpingJimmy2(T, P, O, 1000), 0)
}
