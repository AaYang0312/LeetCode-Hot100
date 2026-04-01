package main

// 「在写递归函数时，可以假设递归返回的结果一定是正确的」
// 自下而上
// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	left := maxDepth(root.Left)
// 	right := maxDepth(root.Right)
// 	return max(left, right) + 1
// }

// 自上而下
func maxDepth(root *TreeNode) (ans int) {
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		depth++
		ans = max(ans, depth)
		dfs(node.Left, depth)
		dfs(node.Right, depth)
	}
	dfs(root, 0)
	return ans
}
