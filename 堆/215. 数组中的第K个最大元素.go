package main

import "math/rand/v2"

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	targetIndex := n - k
	left, right := 0, n-1
	for {
		i := partition(nums, left, right)
		if i == targetIndex {
			return nums[i]
		}
		if i > targetIndex {
			right = i - 1
		} else {
			left = i + 1
		}
	}
}
func partition(nums []int, left, right int) int {
	// 快排
	p := rand.IntN(right-left+1) + left
	pivot := nums[p]
	nums[p], nums[left] = nums[left], nums[p]
	i, j := left+1, right
	for {
		for i <= j && nums[i] < pivot {
			i++
		}
		for i <= j && nums[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	// 循环结束后
	// [ pivot | <=pivot | >=pivot ]
	//   ^             ^   ^     ^
	//   left          j   i     right

	// 3. 把 pivot 与 nums[j] 交换，完成划分（partition）
	// 为什么与 j 交换？
	// 如果与 i 交换，可能会出现 i = right + 1 的情况，已经下标越界了，无法交换
	// 另一个原因是如果 nums[i] > pivot，交换会导致一个大于 pivot 的数出现在子数组最左边，不是有效划分
	// 与 j 交换，即使 j = left，交换也不会出错

	nums[left], nums[j] = nums[j], nums[left]
	return j // pivot 下标
}
