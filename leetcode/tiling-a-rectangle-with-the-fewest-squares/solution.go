package main

import "fmt"

type grid = [13][13]bool

func tilingRectangle(n int, m int) int {
	createGrid := func(c *grid) grid {
		res := grid{}
		if c != nil {
			for i := 0; i < n; i++ {
				for j := 0; j < m; j++ {
					res[i][j] = c[i][j]
				}
			}
		}

		return res
	}

	memo := map[grid]int{}

	var solve func(g grid, depth int) int
	solve = func(g grid, depth int) int {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if g[i][j] {
					fmt.Print(1)
				} else {
					fmt.Print(0)
				}
			}
			fmt.Println()
		}
		fmt.Println()
		fmt.Println()

		if r, f := memo[g]; f {
			return r
		}

		r, f := 13*13+10, true
		if depth > n+m {
			memo[g] = r
			return r
		}

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if !g[i][j] {
					f = false

					/* 	d := n - i - 1
					if m-j-1 < d {
						d = m - j - 1
					}

					for d >= 0 && (g[i+d][j+d] || g[i+d][j] || g[i][j+d]) {
						d--
					} */

					d := 0
					for i+d < n && j+d < m && (!g[i+d][j+d] && !g[i+d][j] && !g[i][j+d]) {
						c := createGrid(&g)
						for k := i; k <= i+d; k++ {
							for l := j; l <= j+d; l++ {
								c[k][l] = true
							}
						}

						next := 1 + solve(c, depth+1)
						if next < r {
							r = next
						}

						d++
					}
				}
			}
		}

		if f {
			r = 0
		}
		memo[g] = r

		return r
	}

	g := createGrid(nil)

	return solve(g, 0)
}

func main() {
	//fmt.Println(3, "==", tilingRectangle(2, 3))
	fmt.Println("?", "==", tilingRectangle(6, 5))
	/* 	fmt.Println(5, "==", tilingRectangle(5, 8))
	   	fmt.Println(6, "==", tilingRectangle(11, 13)) */
}
