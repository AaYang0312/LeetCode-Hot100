package main

func diameterOfBinaryTree(root *TreeNode) int {
	var dfs func(*TreeNode) int
	ans := 0
	dfs = func(root *TreeNode) (depth int) {
		if root == nil {
			return -1 // 叶子结点深度-1,上层-1+1=0
		}
		lLength := dfs(root.Left) + 1
		rLengrh := dfs(root.Right) + 1 // 脚下的深度加上自己

		ans = max(ans, lLength+rLengrh)
		return max(lLength, rLengrh)
	}
	dfs(root)
	return ans
}
