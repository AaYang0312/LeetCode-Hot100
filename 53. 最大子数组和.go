package main

// 思想：动态规划，思考当前数字前面的和是正作用还是副作用
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if max < sum {
			max = sum
		}
	}
	return max
}
