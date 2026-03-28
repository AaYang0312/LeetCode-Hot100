package main

// 封装一个队列
type queue []int

func (q *queue) pop() {
	*q = (*q)[1:len(*q)] // go 左闭右开
}
func (q *queue) push(x int) {
	for len(*q) > 0 && (*q)[len(*q)-1] < x {
		*q = (*q)[:len(*q)-1]
	}
	*q = append(*q, x)
}
func (q queue) getMax() int {
	return q[0]
}
func maxSlidingWindow(nums []int, k int) []int {
	result := []int{}
	q := queue{}
	if len(nums) < k {
		return []int{}
	}
	for i := 0; i < k; i++ {
		q.push(nums[i])
	}
	result = append(result, q.getMax())
	// 推论：一个旧元素如果还在单调队列里，它要么在队头（是最大值），要么早就被删了。它绝不可能安安静静地待在队列中间等着被移除
	for i := k; i < len(nums); i++ {
		if q[0] >= nums[i-k] {
			q.pop()
		}
		q.push(nums[i])
		result = append(result, q.getMax())
	}
	return result
}
