package luckykids

import "fmt"
import "sort"

// Kid is a kid
type Kid struct {
	age, behavior int
}

func (k Kid) String() string {
	return fmt.Sprintf("{ AGE: %v, BEH: %v }\n", k.age, k.behavior)
}

func luckyKids(behaviors []int) int {
	kids := []Kid{}
	L := len(behaviors)

	for i, beh := range behaviors {
		kids = append(kids, Kid{age: i, behavior: beh})
	}

	sort.Slice(kids, func(i, j int) bool {
		if kids[i].behavior == kids[j].behavior {
			return kids[i].age < kids[j].age
		}
		return kids[i].behavior > kids[j].behavior
	})

	naughtyEldersCount := make([]int, L)
	indices := make([]int, L)
	next := make([]int, L)
	streak, b := 0, 0
	for i := 0; i <= L; i++ {
		if i < L && (i == 0 || b == kids[i].behavior) {
			streak++
		} else {
			for j := i - streak; j < i; j++ {
				next[kids[j].age] = i
			}
			streak = 1
		}

		if i == L {
			break
		}
		kid := kids[i]
		next[kid.age] = L

		indices[kid.age] = i
		b = kid.behavior
	}

	top, bottom := 0, L-1
	res := 0
	eldersCountFix := 0
	for _, n := range indices {
		kid := kids[n]
		start := next[kid.age]

		minCount := (L - kid.age) >> 1
		count := eldersCountFix + naughtyEldersCount[n] + L - start

		if count >= minCount {
			res++
		}

		if n <= L>>1 {
			for i := top; i < n; i++ {
				naughtyEldersCount[i]--
			}
			if n == top {
				top++
			}
		} else {
			eldersCountFix--
			//	next := bottom + 1
			for i := n + 1; i <= bottom; i++ {
				//next = i
				naughtyEldersCount[i]++
			}

			if n == bottom {
				bottom--
			}

		}
	}

	return res
}
