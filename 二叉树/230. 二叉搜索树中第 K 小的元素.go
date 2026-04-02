package main

func kthSmallest(root *TreeNode, k int) int {
	var dfs func(*TreeNode)
	i := 1
	num := 0
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		// left
		dfs(node.Left)
		i++
		// mid
		if i == k {
			num = node.Val
			return
		}
		// right
		dfs(node.Right)
	}
	dfs(root)
	return num
}
