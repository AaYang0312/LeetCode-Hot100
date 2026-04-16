package main

func uniquePaths(m int, n int) int {
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 || i >= n || j >= m {
			return 0
		}
		if i == n-1 && j == m-1 {
			return 1
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}

		right := dfs(i+1, j)
		down := dfs(i, j+1)
		memo[i][j] = right + down
		return memo[i][j]
	}
	return dfs(0, 0)
}
