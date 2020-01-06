package main

import "fmt"

func nextPermutation(a []int, l int) ([]int, bool) {
	// Find longest non-increasing suffix
	i := l
	for i > 0 && a[i-1] >= a[i] {
		i--
	}
	// Now i is the head index of the suffix

	// Are we at the last permutation already?
	if i <= 0 {
		return nil, false
	}

	// Let a[i - 1] be the pivot
	// Find rightmost element that exceeds the pivot
	j := len(a) - 1
	d, k := 100, 0
	for j > i-1 {
		t := a[j] - a[i-1]
		if t > 0 && t < d {
			k = j
			d = t
		}
		j--
	}
	j = k

	// Now the value a[j] will become the new pivot
	// Assertion: j >= i

	// Swap the pivot with j
	a[i-1], a[j] = a[j], a[i-1]

	// Reverse the suffix
	j = len(a) - 1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}

	// Successfully computed the next permutation
	return a, true
}

func isSolvable(words []string, result string) bool {
	dict := map[rune]int{}
	abc := []rune{}

	voc := words
	voc = append(voc, result)

	c := 0
	for _, w := range voc {
		for _, ch := range w {
			if _, found := dict[ch]; found {
				continue
			}
			abc = append(abc, rune(ch))

			dict[ch] = c
			c++
		}
	}

	trySolve := func(m map[rune]int) bool {
		sum := 0
		for i, w := range voc {
			if len(w) > 0 && m[rune(w[0])] == 0 {
				return false
			}
			n := 0
			for _, ch := range w {
				n *= 10
				n += m[rune(ch)]
			}

			if i == len(voc)-1 {

				return n == sum
			}
			sum += n
		}

		return true
	}

	p := []int{0, 1, 2, 3, 4, 9, 8, 7, 6, 5}

	for ok := true; ok; {
		fmt.Println(p)
		for i, ch := range abc {
			dict[ch] = p[i]
		}
		if trySolve(dict) {
			return true
		}
		p, ok = nextPermutation(p, len(abc))
	}

	return false
}

func main() {
	fmt.Println(isSolvable([]string{"AA", "BB"}, "DEC")) /*
		fmt.Println(isSolvable([]string{"SEND", "MORE"}, "MONEY"))
		fmt.Println(isSolvable([]string{"THIS", "IS", "TOO"}, "FUNNY"))
		fmt.Println(isSolvable([]string{"LEET", "CODE"}, "POINT")) */
}
