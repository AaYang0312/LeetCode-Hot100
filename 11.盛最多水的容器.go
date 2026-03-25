package main

func maxArea(height []int) int {
	ans := 0
	left, right := 0, len(height)-1
	for left < right {
		area := (right - left) * min(height[left], height[right])
		ans = max(area, ans)
		if height[left] < height[right] {
			left++
		} else if height[left] == height[right] {
			left++
			right--
		} else {
			right--
		}
	}
	return ans
}
