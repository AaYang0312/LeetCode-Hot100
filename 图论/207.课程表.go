package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses) // 存放一门课可开启的下一门课
	for _, p := range prerequisites {
		g[p[1]] = append(g[p[1]], p[0])
	}
	colors := make([]int, numCourses)
	var dfs func(int) bool
	dfs = func(i int) bool {
		colors[i] = 1 // x访问中
		for _, y := range g[i] {
			if colors[y] == 1 {
				return true
			} else if colors[y] == 0 && dfs(y) {
				return true
			}
		}
		colors[i] = 2 // complete
		return false
	}
	for i, c := range colors {
		if c == 0 && dfs(i) {
			return false
		}
	}
	return true
}
