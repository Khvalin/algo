package main

import "fmt"

type statement struct {
	nums string
	hash int
}

func addOperators(num string, target int) []string {
	res := []string{}

	dgts := make([]int, len(num))

	for i, ch := range num {
		dgts[i] = int(ch - '0')
	}

	splitsMemo := map[int][][]int{}

	var nextSplit func(index int) [][]int
	nextSplit = func(index int) [][]int {
		r, found := splitsMemo[index]
		if found {
			return r
		}

		for l := 1; index+l <= len(dgts); l++ {
			n := 0
			for i := index; i-index < l; i++ {
				n *= 10
				n += dgts[i]
			}
			next := nextSplit(index + l)
			if len(next) == 0 {
				r = append(r, []int{n})
			}

			for _, a := range next {
				r = append(r, append([]int{n}, a...))
			}
		}

		splitsMemo[index] = r

		return r
	}

	//opMemo := map[statement]int{}
	combine := func(a []int, k int) (bool, int) {
		ops := make([]int, len(a)-1)
		for i := 0; i < len(a)-1; i++ {
			ops[i] = k % 3
			k /= 3
		}
		if k > 0 {
			return false, 0
		}
		cp := make([]int, len(a))
		copy(cp, a)

		prev, start := 0, -1

		for i, op := range ops {
			if op == 0 {
				cp[i+1] *= cp[i]
				cp[i] = cp[i+1]

				if prev < i-1 {
					start = i
				} else {
					if start >= 0 {
						cp[start] = cp[i]
					}
				}
				prev = i
				if start < 0 {
					start = i
				}
			}
		}

		r := cp[0]
		for i, op := range ops {
			if op > 0 {
				d := -1
				if op == 2 {
					d = 1
				}
				r += d * cp[i+1]
			}
		}

		return true, r
	}

	splits := nextSplit(0)

	for _, split := range splits {
		i := -1

		for valid, r := true, target+1; valid; valid, r = combine(split, i) {
			if r == target {
				s := ""
				//	s += fmt.Sprintf("%v ", r)
				k := i
				for i, n := range split {
					op := ""
					if i > 0 {
						op = string("*-+"[k%3])
						k /= 3
					}
					s += fmt.Sprintf("%v%v", op, n)
				}

				if len(s)-len(split)+1 == len(num) {
					res = append(res, s)
				}
			}
			i++
		}
	}

	return res
}

func main() {
	fmt.Println(addOperators("123", 6))
	fmt.Println(addOperators("0000", 0))
	fmt.Println(addOperators("232", 8))
	fmt.Println(addOperators("105", 5))
	fmt.Println(addOperators("999", 999))
}
