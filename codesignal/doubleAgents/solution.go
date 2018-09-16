package main

import (
	"fmt"
)

func doubleAgents(moles int, reports []bool) string {
	cCount := len(reports) - moles
	cIndices := make([]int, cCount)
	for i := 0; i < cCount; i++ {
		cIndices[i] = i
	}

	L := len(reports)

	res := make([]rune, L)

	overflow := false
	for !overflow {
		cleans := make([]bool, L)

		for i := 0; i < cCount; i++ {
			cleans[cIndices[i]] = true
		}

		//	fmt.Println(cleans)

		valid := true
		for i := 0; valid && (i < L); i++ {
			valid = !cleans[i] || reports[i] == cleans[(i+1)%L]
		}
		if valid {
			for i := 0; i < L; i++ {
				ch := 'C'
				if !cleans[i] {
					ch = 'M'
				}

				if res[i] == 0 {
					res[i] = ch
				} else {
					if res[i] != ch {
						res[i] = '?'
					}
				}
			}
		}

		// next permutation
		i, j := cCount-1, 1
		for i >= 0 && cIndices[i] >= L-j {
			i--
			j++
		}
		overflow = i < 0
		if !overflow {
			cIndices[i]++

			for j := i + 1; j < cCount; j++ {
				cIndices[j] = cIndices[j-1] + 1
			}
			overflow = (cIndices[cCount-1] >= L)
		}

		//	fmt.Println(cIndices, overflow)
	}

	for i, ch := range res {
		if ch == 0 {
			res[i] = '?'
		}
	}

	return string(res)
}

func main() {
	fmt.Println(doubleAgents(1, []bool{true, false, false}))
	fmt.Println(doubleAgents(2, []bool{true, false, false}))
	//	fmt.Println(doubleAgents(3, []bool{true, false, false, true, false, false}))
}
