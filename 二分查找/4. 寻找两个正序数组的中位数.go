package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	nums1 = append([]int{math.MinInt}, append(nums1, math.MaxInt)...)
	nums2 = append([]int{math.MinInt}, append(nums2, math.MaxInt)...)
	left, right := 0, m+1 // right 搜索范围的右边界（不包含在搜索范围内），考虑i=m是答案，如果right取m导致答案被排除
	for left+1 < right {
		i := left + (right-left)/2
		j := (m+n+1)/2 - i
		if nums1[i] <= nums2[j+1] { // nums1[left] <= nums2[j+1]循环不变量
			left = i
		} else {
			right = i
		}
	}
	// left = right - 1
	i := left
	j := (m+n+1)/2 - i
	max1 := max(nums1[i], nums2[j])
	min2 := min(nums1[i+1], nums2[j+1])
	if (m+n)%2 == 0 {
		return (float64(max1 + min2)) / 2.0
	}
	return float64(max1)
}
