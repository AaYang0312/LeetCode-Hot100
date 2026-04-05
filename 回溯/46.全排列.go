package main

func permute(nums []int) [][]int {
	res := [][]int{}                // 储存全排列
	track := []int{}                // 当前选择的数路径
	used := make([]bool, len(nums)) // 标记数字是否被选过
	backtrack(&res, track, used, nums)
	return res
}

func backtrack(res *[][]int, track []int, used []bool, nums []int) {
	// 结束条件
	if len(track) == len(nums) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}
	// 遍历所有可能的选择
	for i := 0; i < len(nums); i++ {
		// 剪枝
		if used[i] {
			continue
		}
		// 选择
		used[i] = true
		track = append(track, nums[i])
		// 递归进入下一层
		backtrack(res, track, used, nums)
		// 递归完成后撤销
		track = track[:len(track)-1]
		used[i] = false
	}
	return
}
