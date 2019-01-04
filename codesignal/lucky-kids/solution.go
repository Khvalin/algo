package luckykids

import "fmt"

import "sort"

// Kid is a kid
type Kid struct {
	age, behavior, naughtyEldersCount int
}

func (k Kid) String() string {
	return fmt.Sprintf("{ AGE: %v, BEH: %v, count: %v }\n", k.age, k.behavior, k.naughtyEldersCount)
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

	res := 0
	eldersCountFix := 0
	for _, n := range indices {
		kid := kids[n]
		start := next[kid.age]

		minCount := (L - kid.age) >> 1
		naughtyEldersCount := eldersCountFix + kid.naughtyEldersCount + L - start

		if naughtyEldersCount >= minCount {
			res++
			//		fmt.Println(naughtyEldersCount, minCount, kid)
		}

		if n < L>>1 {
			for i := 0; i < n; i++ {
				kids[i].naughtyEldersCount--
			}
		} else {
			eldersCountFix--
			for i := n + 1; i < L; i++ {
				kids[i].naughtyEldersCount++
			}
		}
	}
	/*
		fmt.Println(kids)
		fmt.Println(next)
	*/
	return res
}
