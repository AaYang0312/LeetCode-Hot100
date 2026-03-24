package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)
	fmt.Println(nums)
	for n1 := 0; n1 < len(nums)-2; n1++ {
		num1 := nums[n1]
		if num1 > 0 {
			break
		}
		// 跳过重复的n1
		if n1 > 0 && num1 == nums[n1-1] {
			continue
		}
		n2, n3 := n1+1, len(nums)-1
		for n2 < n3 {
			num2, num3 := nums[n2], nums[n3]
			// 暂存三数的和
			sum := num1 + num2 + num3
			// 如果和为零则加入ans中
			if sum == 0 {
				ans = append(ans, []int{num1, num2, num3})
				// 跳过重复的n2
				for n2 < n3 && num2 == nums[n2+1] {
					n2++
				}
				// 跳过重复的n3
				for n2 > n3 && num3 == nums[n3-1] {
					n3--
				}
				n2++
				n3--
			} else if sum < 0 {
				n2++
			} else if sum > 0 {
				n3--
			}
		}
	}
	return ans
}
func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	ans := threeSum(nums)
	fmt.Println(ans)
}
