package leetcode

// MaxProfit2 一个股票最多一股，可以买卖无数次
func MaxProfit2(prices []int) int {
	return greedy(prices)
}

// 贪心
func greedy(prices []int) int {
	var maxProfit int
	for i := 1; i < len(prices); i++ {
		// 只要发现第二天的股票比第一天的低，那么就买入
		profile := prices[i] - prices[i-1]
		if profile > 0 {
			// 买
			maxProfit += profile
		}
	}
	return maxProfit
}
