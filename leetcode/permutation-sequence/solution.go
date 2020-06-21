package main

import "fmt"

func getPermutation(n, k int) string {
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		f[i] = f[i-1] * i
	}

	a := make([]byte, n)
	for i := range a {
		a[i] = '1' + byte(i)
	}

	k--
	j := 0
	res := make([]byte, n)
	for len(a) > 0 {
		i := k / f[n-1]

		res[j] = a[i%n]

		for l := i; l < len(a)-1; l++ {
			a[l] = a[l+1]
		}
		a = a[:n-1]

		k %= f[n-1]
		n--
		j++
	}

	return string(res)
}

func main() {
	var res string
	res = getPermutation(3, 3)
	fmt.Println(res)

	res = getPermutation(4, 9)
	fmt.Println(res)
}
