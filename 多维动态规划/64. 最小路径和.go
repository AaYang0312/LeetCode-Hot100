package main

func minPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = 201
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 || i >= n || j >= m {
			return -1
		}
		if i == n-1 && j == m-1 {
			return grid[i][j]
		}
		if memo[i][j] != 201 {
			return memo[i][j]
		}
		right := dfs(i+1, j)
		down := dfs(i, j+1)
		if right == -1 {
			memo[i][j] = down + grid[i][j]
		} else if down == -1 {
			memo[i][j] = right + grid[i][j]
		} else {
			memo[i][j] = min(right, down) + grid[i][j]
		}
		return memo[i][j]
	}
	return dfs(0, 0)
}
