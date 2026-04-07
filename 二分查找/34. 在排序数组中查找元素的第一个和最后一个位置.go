package main

func searchRange(nums []int, target int) []int {
	start := search(nums, target)
	if start >= len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := search(nums, target+1) - 1
	return []int{start, end}
}

func search(nums []int, target int) int {
	n := len(nums)
	left, right := -1, n
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}
