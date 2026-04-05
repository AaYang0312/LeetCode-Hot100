package main

import "slices"

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	slices.Sort(candidates)
	path := []int{}
	var dfs func(int, int)
	dfs = func(i int, left int) {
		if left == 0 {
			res = append(res, slices.Clone(path))
			return
		}
		// 剪枝优化，candidates 从小到大，如果当前数大于所需值，则不可能
		if i == len(candidates) || left < candidates[i] {
			return
		}
		// 选
		path = append(path, candidates[i])
		dfs(i, left-candidates[i]) // 可以重复，不用加
		path = path[:len(path)-1]
		// 不选
		dfs(i+1, left)
	}
	dfs(0, target)
	return res
}
