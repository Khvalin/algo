package main

func differentSquares(matrix [][]int) int {
	var M map[[4]int]bool
	for i := 0; i < len(matrix)-2; i++ {
		for j := 0; j < len(matrix[i])-2; j++ {
			arr := [4]int{matrix[i][j], matrix[i][j+1], matrix[i+1][j], matrix[i+1][j+1]}
			M[arr] = true
		}
	}

	res := len(M)
	return res
}
