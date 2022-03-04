package leetcode

import "math"

// MaxProfit 只能买一次、卖一次，计算最大利润
func MaxProfit(prices []int) int {
	//return violence(prices)
	return simple(prices)
}

// 暴力破解
func violence(prices []int) int {
	result := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > result {
				result = profit
			}
		}
	}
	return result
}

func simple(prices []int) int {
	minPrice := math.MaxInt
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			// 最低股价
			minPrice = prices[i]
		}

		profit := prices[i] - minPrice
		if profit > maxProfit {
			// 最大利润
			maxProfit = profit
		}
	}
	return maxProfit
}
