package main

func flatten(root *TreeNode) {
	var dfs func(*TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		// left
		left := dfs(node.Left) // left tail
		// right
		right := dfs(node.Right) // right tail

		if left != nil {
			left.Right = node.Right
			node.Right = node.Left // 把原先左子树头节点移到右边
			node.Left = nil        // 将左子树置为空
		}
		if right != nil {
			return right
		}
		if left != nil {
			return left
		}
		return node
	}
	dfs(root)
}
