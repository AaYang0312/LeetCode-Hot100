package main

import "math"

func maxProfit(prices []int) int {
	ans := 0
	minPrice := math.MaxInt
	for _, price := range prices {
		minPrice = min(minPrice, price)
		ans = max(ans, price-minPrice)
	}
	return ans
}
