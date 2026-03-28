package main

func rotate(nums []int, k int) {
	m := len(nums)
	if m == 1 || k%m == 0 {
		return
	}
	k = k % m
	nums1 := nums[0 : m-k]
	nums2 := nums[m-k : m]
	rotated := append(nums2, nums1...)
	copy(nums, rotated) // append时发现容量不够，扩容会新开一块内存块，必须使用copy赋值
}
