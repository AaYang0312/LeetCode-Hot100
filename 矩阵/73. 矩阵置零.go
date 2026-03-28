package main

func setZeroes(matrix [][]int) {
	row0, col0 := false, false
	m, n := len(matrix[0]), len(matrix)
	// 用两个变量记录第0行和第0列是否有0，再第0行、第0列记录其他行列是否有0
	for _, v := range matrix[0] {
		if v == 0 {
			row0 = true
			break
		}
	}
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
			break
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if row0 {
		for i := 0; i < m; i++ {
			matrix[0][i] = 0
		}

	}
	if col0 {
		for j := 0; j < n; j++ {
			matrix[j][0] = 0
		}
	}
}
