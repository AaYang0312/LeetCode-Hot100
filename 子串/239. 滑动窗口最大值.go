func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}
	ans := make([]int, len(nums)-k+1)
	q := []int{}
	for i, x := range nums {
		// 右边进
		for len(q) > 0 && nums[q[len(q)-1]] <= x {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		// 左边出
		left := i - k + 1
		if left > q[0] {
			q = q[1:]
		}
		// 记录答案
		if left >= 0 {
			ans[left] = nums[q[0]]
		}
	}
	return ans
}
