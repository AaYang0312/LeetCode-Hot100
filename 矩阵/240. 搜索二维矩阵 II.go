package main

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix[0]), len(matrix)
	i, j := 0, m-1
	for i < n && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}
