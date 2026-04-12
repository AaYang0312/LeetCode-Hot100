package main

func rob(nums []int) int {
	f0, f1 := 0, 0 // i-2, i-1 个房子的最大值, (i+1)-2=i-1，所以下一轮的f0=f1

	for _, x := range nums {
		f0, f1 = f1, max(f0+x, f1)
	}
	return f1
}
