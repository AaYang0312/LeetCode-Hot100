package main

// func productExceptSelf(nums []int) []int {
// 	m := len(nums)
// 	L, R, ans := make([]int, m), make([]int, m), make([]int, m)
// 	L[0] = 1
// 	for i := 1; i < m; i++ {
// 		L[i] = L[i-1] * nums[i-1]
// 	}
// 	R[m-1] = 1
// 	for i := m - 2; i >= 0; i-- {
// 		R[i] = R[i+1] * nums[i+1]
// 	}
// 	for i := 0; i < m; i++ {
// 		ans[i] = L[i] * R[i]
// 	}
// 	return ans
// }
func productExceptSelf(nums []int) []int {
	m := len(nums)
	ans := make([]int, m)
	ans[0] = 1
	for i := 1; i < m; i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	R := 1
	for i := m - 1; i >= 0; i-- {
		ans[i] = ans[i] * R
		R *= nums[i]
	}
	return ans
}
