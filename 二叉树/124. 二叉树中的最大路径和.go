package main

import "math"

func maxPathSum(root *TreeNode) int {
	var dfs func(*TreeNode) int
	var maxP = math.MinInt
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftMax := dfs(node.Left)
		rightMax := dfs(node.Right)
		maxP = max(leftMax+rightMax+node.Val, maxP)
		return max(max(leftMax, rightMax)+node.Val, 0)
	}
	dfs(root)
	return maxP
}
