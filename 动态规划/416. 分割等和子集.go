package main

func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 奇数的一半不可能是整数
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2
	visited := make([][]bool, n)
	memo := make([][]bool, n)
	for i := range memo {
		visited[i] = make([]bool, sum+1)
		memo[i] = make([]bool, sum+1)
	}
	var dfs func(int, int) bool
	dfs = func(i, j int) bool {
		if i < 0 {
			return j == 0
		}
		if visited[i][j] {
			return memo[i][j]
		}
		res := j >= nums[i] && dfs(i-1, j-nums[i]) || dfs(i-1, j)
		memo[i][j] = res
		visited[i][j] = true
		return res
	}
	return dfs(n-1, sum)
}
