package main

import "fmt"

func gridIllumination(N int, lamps [][]int, queries [][]int) []int {
	type INT = uint32
	type set = map[INT]bool

	const MaxInt = 20001

	type coord struct {
		x INT
		y INT
	}

	ranks := map[INT]set{}
	files := map[INT]set{}
	diags := map[coord]set{}

	normalizeDiag := func(x, y INT) (coord, coord) {
		return coord{x: x - y}, coord{y: x + y}
	}

	grid := map[coord]INT{}

	//	lamps := make([]coord, len(lampsArr))
	for i, lamp := range lamps {
		x, y := INT(lamp[0]), INT(lamp[1])
		index := INT(i)
		//lamps[i] = coord{x: x, y: y}

		if _, ok := ranks[x]; !ok {
			ranks[x] = set{}
		}

		ranks[x][index] = true

		if _, ok := files[y]; !ok {
			files[y] = set{}
		}

		files[y][index] = true

		dNeg, dPos := normalizeDiag(x, y)
		if _, ok := diags[dNeg]; !ok {
			diags[dNeg] = set{}
		}
		diags[dNeg][index] = true

		if _, ok := diags[dPos]; !ok {
			diags[dPos] = set{}
		}
		diags[dPos][index] = true

		grid[coord{x, y}] = index
	}

	res := make([]int, len(queries))
	for i, q := range queries {
		r := false
		x, y := INT(q[0]), INT(q[1])

		if rank, ok := ranks[x]; ok {
			r = r || len(rank) > 0
		}

		if file, ok := files[y]; ok {
			r = r || len(file) > 0
		}

		dNeg, dPos := normalizeDiag(x, y)

		if diag, ok := diags[dNeg]; ok {
			r = r || len(diag) > 0
		}

		if diag, ok := diags[dPos]; ok {
			r = r || len(diag) > 0
		}

		k := x - 1
		if k < 0 {
			k = 0
		}

		l := y - 1
		if l < 0 {
			l = 0
		}

		kMax := x + 1
		if kMax > INT(N-1) {
			kMax = INT(N - 1)
		}

		lMax := l + 1
		if lMax > INT(N-1) {
			lMax = INT(N - 1)
		}

		for k <= kMax {
			for l <= lMax {
				c := coord{k, l}
				if index, ok := grid[c]; ok {
					dNeg, dPos := normalizeDiag(k, l)

					delete(ranks[k], index)
					delete(files[l], index)
					delete(diags[dNeg], index)
					delete(diags[dPos], index)

					delete(grid, c)
				}
				l++
			}
			k++

		}

		if r {
			res[i] = 1
		}
	}

	return res
}

func main() {
	t1 := gridIllumination(5, [][]int{[]int{0, 0}, []int{4, 4}}, [][]int{[]int{1, 1}, []int{1, 0}})
	fmt.Println(t1, uint32(1000000000))

}
