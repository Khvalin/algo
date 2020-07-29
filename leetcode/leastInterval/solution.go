package main

import "fmt"

func leastInterval(tasks []byte, n int) int {
	count := make([]int, 'Z')
	last := make([]int, 'Z')
	abc := []byte{}

	for _, ch := range tasks {
		if count[ch] == 0 {
			abc = append(abc, ch)
		}
		count[ch]++
		last[ch] = -n - 1
	}

	res := 0
	for _ = range tasks {
		next, min := byte(0), n+10

		for _, ch := range abc {
			score := n - (res - last[ch]) + 1
			if score < 0 {
				score = 0
			}

			if count[ch] > 0 && (score < min || score == min && count[ch] > count[next]) {
				next = ch
				min = score
			}
		}

		res += min
		last[next] = res
		count[next]--

		res++
	}

	return res
}

func main() {
	r := leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2)
	fmt.Println()
	fmt.Println(r, 8)

	r = leastInterval([]byte{'A', 'A'}, 2)
	fmt.Println()
	fmt.Println(r, 4)

	r = leastInterval([]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2)
	fmt.Println()
	fmt.Print(r, 16)
}
