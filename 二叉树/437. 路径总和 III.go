package main

func pathSum(root *TreeNode, targetSum int) (ans int) {
	cnt := map[int]int{0: 1}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, k int) {
		if node == nil {
			return
		}
		k += node.Val           // 当前根到节点和
		ans += cnt[k-targetSum] // 查找有几个前缀使 k-前缀=目标值
		cnt[k]++

		dfs(node.Left, k)
		dfs(node.Right, k)
		cnt[k]-- // 向上退时去掉当前节点
	}
	dfs(root, 0)
	return ans
}
