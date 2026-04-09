package main

func jump(nums []int) int {
	nextRight := 0
	currRight := 0
	cnt := 0
	for i, jump := range nums[:len(nums)-1] {
		nextRight = max(nextRight, i+jump)
		if i == currRight {
			cnt++
			currRight = nextRight
		}
	}
	return cnt
}
