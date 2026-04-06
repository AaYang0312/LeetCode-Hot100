package main

import "slices"

func partition(s string) [][]string {
	ans := [][]string{}
	n := len(s)
	path := []string{}
	var dfs func(int, int)
	dfs = func(end int, start int) {
		if end == n {
			ans = append(ans, slices.Clone(path))
			return
		}
		// 不切
		if end < n-1 {
			dfs(end+1, start)
		}
		// 切
		substr := s[start : end+1]
		if isPalindrome(substr) {
			path = append(path, substr)
			dfs(end+1, end+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return ans
}

func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}
