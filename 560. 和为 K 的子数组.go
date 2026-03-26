package main

func subarraySum(nums []int, k int) int {
	hash := map[int]int{0: 1} // 方便统计第一个元素=k情况
	count, preSum := 0, 0
	for i := 0; i < len(nums); i++ {

		// 首先求当前前缀和
		preSum += nums[i]

		// 判断是否存在历史前缀和，使 当前-历史=k（即存在一个子数组） 成立
		if hash[preSum-k] > 0 {
			// 有多少个前缀，当前就多出多少不同子数组，所以累加
			count += hash[preSum-k]
		}

		// 当前前缀和出现次数纳入统计
		hash[preSum]++
	}
	return count
}
