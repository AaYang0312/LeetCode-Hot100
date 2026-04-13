package main

import "sort"

//import "slices"
//
//func lengthOfLIS(nums []int) int {
//	n := len(nums)
//	numsMap := make([]int, n)
//	for i := 0; i < n; i++ {
//		numsMap[i] = 1
//	}
//	for i := 0; i < n; i++ {
//		for j := 0; j < i; j++ {
//			if nums[j] < nums[i] {
//				numsMap[i] = max(numsMap[i], numsMap[j]+1)
//			}
//		}
//	}
//	return slices.Max(numsMap)
//}

func lengthOfLIS(nums []int) int {
	g := []int{}
	for _, x := range nums {
		j := sort.SearchInts(g, x)
		if j == len(g) { // >=x 的 g[j] 不存在
			g = append(g, x)
		} else {
			g[j] = x
		}
	}
	return len(g)
}
