package main

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left+right)/2 + left
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}
