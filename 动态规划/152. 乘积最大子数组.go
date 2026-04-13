package main

import "math"

func maxProduct(nums []int) int {
	fmax, fmin := 1, 1
	ans := math.MinInt
	for _, x := range nums {
		fmax, fmin = max(fmax*x, fmin*x, x), min(fmin*x, fmax*x, x)
		ans = max(ans, fmax)
	}
	return ans
}
