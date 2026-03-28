package main

func longestConsecutive(nums []int) int {
	// 定义一个哈希表去重
	theHash := make(map[int]bool)
	// 储存最大长度
	result := 0
	// 首先将int切片读入哈希表
	for _, num := range nums {
		theHash[num] = true
	}
	// 当前数字
	// currNum := 0
	// 当前连续序列长度
	// currLen := 0
	// 遍历哈希表
	for num, _ := range theHash {
		if theHash[num-1] {
			continue
		}
		currNum := num
		currLen := 1
		for theHash[currNum+1] {
			currNum++
			currLen++
		}
		if result < currLen {
			result = currLen
		}
	}
	return result
}
