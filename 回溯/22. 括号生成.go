package main

func generateParenthesis(n int) []string {
	res := []string{}
	path := make([]byte, n*2)
	var dfs func(int, int)
	dfs = func(left int, right int) {
		if left+right == 2*n {
			res = append(res, string(path))
			return
		}
		if left < n {
			path[left+right] = '('
			dfs(left+1, right)
		}
		if right < left {
			path[left+right] = ')'
			dfs(left, right+1)
		}
	}
	dfs(0, 0)
	return res
}
