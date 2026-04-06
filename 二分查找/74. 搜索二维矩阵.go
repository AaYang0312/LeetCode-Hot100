package main

func searchMatrix(matrix [][]int, target int) bool {
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	for top <= bottom {
		midY := (bottom-top)/2 + top
		if matrix[midY][0] < target {
			top = midY + 1
		} else if matrix[midY][0] > target {
			bottom = midY - 1
		} else {
			return true
		}
	}
	if bottom >= len(matrix) || bottom < 0 {
		return false
	}
	for left <= right {
		midX := (right-left)/2 + left
		if matrix[bottom][midX] < target {
			left = midX + 1
		} else if matrix[bottom][midX] > target {
			right = midX - 1
		} else {
			return true
		}
	}
	return false
}
