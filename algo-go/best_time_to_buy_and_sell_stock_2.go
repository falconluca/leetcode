package leetcode

// MaxProfit2 一个股票最多一股，可以买卖无数次
func MaxProfit2(prices []int) int {
	//return greedy(prices)
	return dp(prices)
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

func dp(prices []int) int {
	n := len(prices)
	// 第i天的最大利润, 是否持有(0/1)
	maxProfits := make([][2]int, n)

	maxProfits[0][0] = 0
	maxProfits[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		maxProfits[i][0] = max(maxProfits[i-1][0], // 不动
			maxProfits[i-1][1]+prices[i]) // 卖出
		maxProfits[i][1] = max(maxProfits[i-1][1], // 不动
			maxProfits[i-1][0]-prices[i]) // 买入
	}

	return maxProfits[n-1][0]
}
