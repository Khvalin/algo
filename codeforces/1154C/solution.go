package main

import "fmt"

func solve(a [3]uint64) uint64 {
	week := [...]int{0, 1, 2, 0, 2, 1, 0}
	for i := 0; i < 7; i++ {
		v := [3]uint64{}
		for j := 0; j < 7; j++ {
			d := (i + j) % 7
			v[week[d]]++
		}
		fmt.Println(v)
	}

	return 0
}

func main() {
	a := [3]uint64{}

	for i := 0; i < 3; i++ {
		fmt.Scan(&a[i])
	}

	ans := solve(a)

	fmt.Println(ans)
}
