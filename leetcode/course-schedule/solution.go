package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	connections := make([][]int, numCourses)
	count := make([]int, numCourses)
	for _, p := range prerequisites {
		connections[p[1]] = append(connections[p[1]], p[0])
		count[p[0]]++
	}

	total := 0
	a := make([]int, 0, numCourses)
	for i := 0; i < numCourses; i++ {
		if count[i] == 0 {
			a = append(a, i)
			total++
		}
	}

	for total < numCourses && len(a) > 0 {
		l := len(a)
		for i := 0; i < l; i++ {
			ind := a[i]
			for _, j := range connections[ind] {
				count[j]--
				if count[j] == 0 {
					a = append(a, j)
					total++
				}
			}
		}

		a = a[l:]
	}

	return total == numCourses
}

func main() {
	fmt.Println(canFinish(2, [][]int{{0, 1}}))
}
