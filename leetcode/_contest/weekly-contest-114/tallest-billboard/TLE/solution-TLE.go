package tallestBillboard

import "fmt"

func tallestBillboard(rods []int) int {
	N := len(rods)

	sums := map[int][]int{}
	res := 0
	arr := make([]int, N+1)
	arr[0] = 1
	for arr[N] == 0 {
		acc, s := 0, 0

		for i := N - 1; i >= 0; i-- {
			acc *= 2
			acc += arr[i]

			if arr[i] == 1 {
				s += rods[i]
			}
		}
		val, found := sums[s]
		if !found {
			val = []int{}
		}
		val = append(val, acc)
		sums[s] = val

		// add one:
		i := 0
		for arr[i] == 1 {
			arr[i] = 0
			i++
		}
		arr[i]++
	}

	for k, v := range sums {
		if len(v) > 1 && k > res {
			//	fmt.Println(k, v)
			for i := 0; i < len(v); i++ {
				for j := i + 1; j < len(v); j++ {
					if v[i]&v[j] == 0 {
						res = k
						break
					}
				}
				if res == k {
					break
				}
			}
		}
	}

	return res
}

func main() {

	fmt.Println(tallestBillboard([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(tallestBillboard([]int{1, 2, 3, 6}))
	fmt.Println(tallestBillboard([]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 50, 50, 50, 150, 150, 150, 100, 100, 100, 123}))

	fmt.Println(tallestBillboard([]int{124, 121, 107, 127, 156, 146, 135, 153, 137, 150, 141, 138, 129, 142, 124, 144, 126, 900, 900, 900}))
}
