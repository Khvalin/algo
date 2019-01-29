package main

import "fmt"

func productFromDifferentDigits(n int) (r int) {
	j := -1
	for i := 0; j < n && i < 362881; i++ {
		dd := [...]int{6, 8, 9, 4, 3, 2, 5, 7}
		k, d := i, 9
		for _, d = range dd {
			if k%d == 0 {
				k /= d
			}
			d--
		}
		if k < 2 {
			r = i
			j++
		}
	}

	//	fmt.Println(j, r)

	if j < n {
		r = -1
	}
	return
}

func main() {
	for i := 0; i < 150; i++ {
		fmt.Println(i, productFromDifferentDigits(i))
	}

	fmt.Println(648, productFromDifferentDigits(78))
	fmt.Println(productFromDifferentDigits(151))
	fmt.Println(productFromDifferentDigits(152))
}
