package main

import "fmt"

func minFlips(mat [][]int) int {
	n, m := len(mat), len(mat[0])

	c := [3][3]bool{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			c[i][j] = 1 == mat[i][j]
		}
	}

	steps := map[[3][3]bool]int{}

	var calc func(mat [3][3]bool, s int)
	calc = func(mat [3][3]bool, s int) {
		if v, ok := steps[mat]; v <= s && ok {
			return
		}

		steps[mat] = s

		f := true
		c := [3][3]bool{}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				c[i][j] = mat[i][j]
				f = f && !c[i][j]
			}
		}

		flip := func(c *[3][3]bool, i, j int) {
			if i >= 0 && i < n && j >= 0 && j < m {
				c[i][j] = !c[i][j]
			}
		}

		if !f {
			for i := 0; i < n; i++ {
				for j := 0; j < m; j++ {
					cc := [3][3]bool{}

					cc = c

					flip(&cc, i, j)
					flip(&cc, i+1, j)
					flip(&cc, i-1, j)
					flip(&cc, i, j-1)
					flip(&cc, i, j+1)

					calc(cc, s+1)
				}
			}
		}
	}

	calc(c, 0)
	zeroes := [3][3]bool{}

	r, ok := steps[zeroes]
	if !ok {
		return -1
	}
	return r
}

func main() {
	mat := [][]int{{0, 0}, {0, 1}}
	fmt.Println(3, minFlips(mat))
	//mat = [][]int{{1, 1}, {0, 1}}
	//mat = [][]int{{1, 1, 1}, {1, 0, 1}, {0, 0, 0}}
	//mat = [][]int{{1, 0, 0}, {1, 0, 0}}
	//	mat = [][]int{{0}, {1}, {1}}
	mat = [][]int{{1, 1, 1}, {1, 0, 1}, {0, 0, 0}}
	fmt.Println(6, minFlips(mat))

	//fmt.Println(6, minFlips(mat))

}
