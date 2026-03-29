package main

func spiralOrder(matrix [][]int) []int {
	ans := []int{}
	m, n := len(matrix[0]), len(matrix)
	top, bottom, left, right := 0, n-1, 0, m-1

	for left <= right && top <= bottom {
		// 右
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
		top++
		// 下
		for i := top; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--
		// 左
		if top <= bottom {
			for i := right; i >= left; i-- {
				ans = append(ans, matrix[bottom][i])
			}
			bottom--
		}

		// 上
		if left <= right {
			for i := bottom; i >= top; i-- {
				ans = append(ans, matrix[i][left])
			}
			left++
		}
	}
	return ans
}
