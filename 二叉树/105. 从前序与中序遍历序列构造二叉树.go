package main

//	func buildTree(preorder, inorder []int) *TreeNode {
//		n := len(preorder)
//		if n == 0 { // 空节点
//			return nil
//		}
//		leftSize := slices.Index(inorder, preorder[0]) // 左子树的大小
//		left := buildTree(preorder[1:1+leftSize], inorder[:leftSize])
//		// 左子树的所有节点，共 leftSize 个
//		right := buildTree(preorder[1+leftSize:], inorder[1+leftSize:])
//		return &TreeNode{preorder[0], left, right}
//	}
func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	InIndex := make(map[int]int, n)
	for i, x := range inorder {
		InIndex[x] = i
	}
	var dfs func(int, int, int) *TreeNode
	dfs = func(preL, preR, inL int) *TreeNode {
		if preL == preR {
			return nil
		}
		leftSize := InIndex[preorder[preL]] - inL
		left := dfs(preL+1, preL+1+leftSize, inL)
		right := dfs(preL+1+leftSize, preR, inL+1+leftSize)
		return &TreeNode{preorder[preL], left, right}
	}
	return dfs(0, n, 0) // 左闭右开区间
}
