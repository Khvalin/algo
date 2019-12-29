package main

func canReach(arr []int, start int) bool {
	visited := make([]bool, len(arr)+1)
	q := []int{start}
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		if visited[i] {
			continue
		}
		visited[i] = true

		if arr[i] == 0 {
			return true
		}

		l, r := i-arr[i], i+arr[i]
		if l >= 0 && l < len(arr) && !visited[l] {
			q = append(q, l)
		}

		if r >= 0 && r < len(arr) && !visited[r] {
			q = append(q, r)
		}
	}

	return false
}

func main() {}
