package main

import "math"

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}
	pre := math.MinInt
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		// 左
		if !dfs(node.Left) {
			return false
		}
		// 中
		if pre >= node.Val {
			return false
		}
		// 右
		pre = node.Val
		return dfs(node.Right)
	}
	return dfs(root)
}
