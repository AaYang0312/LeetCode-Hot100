package main

import "strings"

func solveNQueens(n int) [][]string {
	ans := [][]string{}
	queens := make([]int, n)
	// 存储对角线，一条上x+y相等，一条上x-y相等
	diag1 := make([]bool, 2*n-1)
	diag2 := make([]bool, 2*n-1)
	col := make([]bool, n) // 每列最多一个
	var dfs func(int)
	dfs = func(r int) {
		if r == n {
			board := make([]string, n)
			for i, c := range queens {
				board[i] = strings.Repeat(".", c) + "Q" + strings.Repeat(".", n-c-1)
			}
			ans = append(ans, board)
			return
		}
		// 放在(r, l)
		for l, ok := range col {
			if !ok && !diag1[r+l] && !diag2[r-l+n-1] {
				queens[r] = l // 存第r个queen在第l列
				col[l], diag1[r+l], diag2[r-l+n-1] = true, true, true
				dfs(r + 1)
				col[l], diag1[r+l], diag2[r-l+n-1] = false, false, false
			}
		}
	}
	dfs(0)
	return ans
}
