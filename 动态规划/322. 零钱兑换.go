package main

func coinChange(coins []int, amount int) int {
	res := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		res[i] = amount + 1
	}
	// 遍历物品
	for _, coin := range coins {
		// 遍历背包容量
		for j := coin; j <= amount; j++ {
			res[j] = min(res[j], res[j-coin]+1)
		}
	}
	if res[amount] > amount {
		return -1
	}
	return res[amount]
}
