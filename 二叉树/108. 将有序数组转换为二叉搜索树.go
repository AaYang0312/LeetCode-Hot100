package main

func sortedArrayToBST(nums []int) *TreeNode {
	m := len(nums)
	if m == 0 {
		return nil
	}
	return &TreeNode{
		Val:   nums[m/2],
		Left:  sortedArrayToBST(nums[:m/2]),
		Right: sortedArrayToBST(nums[m/2+1:]),
	}
}
