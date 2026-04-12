package main

import "math"

const N = 10000

var f [N + 1]int

func init() {
	// 初始化
	for i := 1; i <= N; i++ {
		f[i] = math.MaxInt
	}
	for i := 1; i <= N; i++ {
		// 完全背包，正序遍历
		for j := i * i; j <= N; j++ {
			f[j] = min(f[j], f[j-i*i]+1)
		}
	}
}
func numSquares(n int) int {
	return f[n]
}
