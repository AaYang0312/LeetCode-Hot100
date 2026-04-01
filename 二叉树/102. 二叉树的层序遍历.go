package main

func levelOrder(root *TreeNode) (nodes [][]int) {
	if root == nil {
		return nil
	}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, i int) {
		if node == nil {
			return
		}
		// 确保有足够层数
		if i == len(nodes) {
			nodes = append(nodes, []int{})
		}
		nodes[i] = append(nodes[i], node.Val)
		dfs(node.Left, i+1)
		dfs(node.Right, i+1)
	}
	dfs(root, 0)
	return
}
