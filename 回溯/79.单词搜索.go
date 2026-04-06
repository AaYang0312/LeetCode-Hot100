package main

import "slices"

var dir = []struct{ x, y int }{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func exist(board [][]byte, word string) bool {
	cnt := map[byte]int{}
	for _, row := range board {
		for _, c := range row {
			cnt[c]++
		}
	}

	// 优化一
	w := []byte(word)
	wordCnt := map[byte]int{}
	for _, c := range w {
		wordCnt[c]++
		if wordCnt[c] > cnt[c] {
			return false
		}
	}

	// 优化二
	if cnt[w[len(w)-1]] < cnt[w[0]] {
		slices.Reverse(w)
	}

	m, n := len(board), len(board[0])
	var dfs func(int, int, int) bool
	dfs = func(i int, j int, k int) bool {
		// fail
		if board[i][j] != word[k] {
			return false
		}
		// success
		if k == len(word)-1 {
			return true
		}
		// 标记是否用过
		board[i][j] = 0
		for _, d := range dir {
			x := i + d.x
			y := j + d.y
			if x < 0 || y < 0 || x >= m || y >= n || board[x][y] == 0 {
				continue
			}
			if dfs(x, y, k+1) {
				return true
			}
		}
		board[i][j] = word[k]
		return false

	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
