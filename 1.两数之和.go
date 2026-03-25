package main

func twoSum(nums []int, target int) []int {
	prevNums := map[int]int{}
	temp := nums
	for i, num := range temp {
		targetNum := target - num
		index, ok := prevNums[targetNum]
		if ok {
			return []int{index, i}
		} else {
			prevNums[num] = i
		}
	}
	return []int{}
}
