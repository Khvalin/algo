package main

import (
	"fmt"
)

func solve(n, k int) (bool, []int) {
	var POW []int

	if k > n {
		return false, nil
	}

	p := (1)
	for p <= n {
		POW = append(POW, p)
		p *= 2
	}

	var res []int

	i := len(POW) - 1
	for n > 0 {
		if n >= POW[i] {
			res = append(res, POW[i])
			n -= POW[i]
		}
		i--
	}

	if k < len(res) {
		return false, nil
	}

	for len(res) < k {
		fmt.Println(res)

		if res[0] == 1 {
			return false, nil
		}

		if res[0] == 2 {
			res = res[1:]
			res = append(res, 1)
			res = append(res, 1)
			continue
		}

		res[0] /= 2
		res = append([]int{res[0]}, res...)
	}

	return true, res
}

func main() {
	var n, k int

	fmt.Scan(&n, &k)

	isPossible, a := (solve(n, k))
	if !isPossible {
		fmt.Println("NO")
		return
	}

	fmt.Println("YES")
	for _, n := range a {
		fmt.Print(n, " ")
	}
	fmt.Println()
}
