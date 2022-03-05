package lintcode

import "math"

// 2 5 7
// 27
// 最少硬币付清

// 递归
// 递归+缓存
// 数组dp
// 变量dp

// You are given coins of different denominations and a total amount of money amount.
// Write a function to compute the fewest number of coins that you need to make up that amount.
// If that amount of money cannot be made up by any combination of the coins, return -1.

func CoinChange(coins []int, amount int) int {
	// 0...amount: dp[amount+1]
	// 0...amount-1: dp[amount]
	dp := make([]int, amount+1)

	dp[0] = 0
	for i := 1; i <= amount; i++ { // amount 而不是 len(dp)
		dp[i] = math.MaxInt // 表示无法拼成

		for j := 0; j < len(coins); j++ {
			//nums := dp[i-coins[j]]
			if i >= coins[j] && dp[i-coins[j]] != math.MaxInt {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt {
		dp[amount] = -1
	}
	return dp[amount]
}

func dfs(n int) int {
	res := math.MaxInt
	if n == 0 {
		return 0
	}
	if n >= 2 {
		res = min(dfs(n-2)+1, res)
	}
	if n >= 5 {
		res = min(dfs(n-5)+1, res)
	}
	if n >= 7 {
		res = min(dfs(n-7)+1, res)
	}
	return res
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
