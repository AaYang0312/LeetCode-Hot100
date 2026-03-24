package main

import "fmt"

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

func main() {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	m := maxArea(a)
	fmt.Println(m)
}
