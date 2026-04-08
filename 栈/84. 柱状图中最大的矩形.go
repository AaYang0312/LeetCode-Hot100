package main

func largestRectangleArea(heights []int) int {
	ans := 0
	heights = append(heights, -1) // 最后强制弹出栈中所有元素
	st := []int{-1}               // 对应栈最左边left = -1
	for right, height := range heights {
		for len(st) > 1 && height <= heights[st[len(st)-1]] {
			i := st[len(st)-1]
			h := heights[i]
			st = st[:len(st)-1]
			left := st[len(st)-1]
			ans = max(ans, h*(right-left-1))
		}
		st = append(st, right)
	}
	return ans
}
