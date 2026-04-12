package main

func generate(numRows int) [][]int {
	nums := make([][]int, numRows)
	for i := 1; i <= numRows; i++ {
		start, end := 0, i-1
		nums[i][start], nums[i][end] = 1, 1
		p := start + 1
		for p < end {
			p := start + 1
			nums[i][p] = nums[i-1][p-1] + nums[i-1][p]
			p++
		}
	}
	return nums
}
