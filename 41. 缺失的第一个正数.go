package main

func firstMissingPositive(nums []int) int {
	m := len(nums)

	// 标记所有负数,以及0
	for i, num := range nums {
		if num <= 0 {
			nums[i] = m + 1
		}
	}
	// 用负号标记所有(1,m)间的整数
	// 需要取绝对值，因为可能被前面的操作添加负号，例如处理[1,1]
	for i := 0; i < m; i++ {
		numa := abs(nums[i])
		if numa >= 1 && numa <= m {
			if nums[numa-1] > 0 {
				nums[numa-1] = -nums[numa-1]
			}
		}

	}
	for i := 0; i < m; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return m + 1
}
func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}
