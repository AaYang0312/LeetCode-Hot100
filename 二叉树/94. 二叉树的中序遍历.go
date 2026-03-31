package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func inorderTraversal(root *TreeNode) (ans []int) {
// 	var dfs func(*TreeNode)
// 	dfs = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		dfs(node.Left)              // 左
// 		ans = append(ans, node.Val) // 根（这行代码移到前面就是前序，移到后面就是后序）
// 		dfs(node.Right)             // 右
// 	}
// 	dfs(root)
// 	return
// }

// 不用递归
func inorderTraversal(root *TreeNode) (ans []int) {
	for root != nil {
		if root.Left != nil {
			// 找 root 的前驱 pre：在中序遍历中，root 的上一个节点
			// 从 root.Left 开始，一直向右走，直到走到尽头，或者遇到指向 root 的线索（回到 root 的路）
			pre := root.Left
			for pre.Right != nil && pre.Right != root {
				pre = pre.Right
			}
			// root 的左子树尚未访问
			if pre.Right == nil {
				pre.Right = root // 建立索引
				root = root.Left
				continue
			}
			// root 的左子树访问完毕，复原
			pre.Right = nil
		}
		ans = append(ans, root.Val)
		root = root.Right // 如果有右子树就访问右子树，没有就顺着线索回到指向的节点
	}
	return ans
}
