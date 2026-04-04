package main

type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	cnt := 0
	queue := []pair{}
	for i, row := range grid {
		for j, orange := range row {
			if orange == 1 {
				cnt++ // 统计新鲜橘子数量
			} else if orange == 2 {
				queue = append(queue, pair{i, j})
			}
		}
	}
	time := 0
	for cnt > 0 && len(queue) > 0 {
		time++
		tmp := queue
		queue = []pair{}
		for _, p := range tmp {
			for _, q := range directions {
				i := p.x + q.x
				j := p.y + q.y
				if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 0 || grid[i][j] == 2 {
					continue
				}
				cnt--
				grid[i][j] = 2
				queue = append(queue, pair{i, j})
			}
		}
	}
	if cnt > 0 {
		return -1
	}
	return time
}
