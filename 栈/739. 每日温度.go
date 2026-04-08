package main

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	st := []int{} // 记录还没算出下一个更大元素的那些数的下标
	for i, t := range temperatures {
		for len(st) > 0 && t > temperatures[st[len(st)-1]] {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			ans[j] = i - j
		}
		st = append(st, i)
	}
	return ans
}
