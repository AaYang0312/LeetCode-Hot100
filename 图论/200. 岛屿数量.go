package main

func numIslands(grid [][]byte) int {
	var dfs func(int, int)
	//  递归是递归一座岛
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '2'
		// left
		dfs(i, j-1)
		// right
		dfs(i, j+1)
		// up
		dfs(i-1, j)
		// down
		dfs(i+1, j)
	}
	cnt := 0
	for i, row := range grid {
		for j, c := range row {
			if c == '1' {
				cnt++
				dfs(i, j) // 把这个岛插满旗子
			}
		}
	}
	return cnt
}
