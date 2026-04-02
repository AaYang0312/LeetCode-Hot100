package main

func rightSideView(root *TreeNode) []int {
	var dfs func(*TreeNode, int)
	nodes := [][]int{}
	dfs = func(node *TreeNode, k int) {
		if node == nil {
			return
		}
		if len(nodes) == k {
			nodes = append(nodes, []int{})
		}
		nodes[k] = append(nodes[k], node.Val)
		dfs(node.Left, k+1)
		dfs(node.Right, k+1)
		return
	}
	dfs(root, 0)
	n := []int{}
	for i := 0; i < len(nodes); i++ {
		n = append(n, nodes[i][len(nodes[i])-1])
	}
	return n
}
