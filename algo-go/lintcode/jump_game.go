package lintcode

// https://www.lintcode.com/problem/116/
//
// Given an array of non-negative integers, you are initially positioned at the first index of the array.
//
// Each element in the array represents your maximum jump length at that position.
//
// Determine if you are able to reach the last index.

func CanJump(stones []int) bool {
	nums := len(stones)
	dp := make([]bool, nums)
	dp[0] = true

	for i := 1; i < nums; i++ {
		// 初始化，假设从当前位置无法到达石头 i
		dp[i] = false

		// 前面已经到达的石头，是否有可以到达石头 i 的
		for j := 0; j < i; j++ {
			// stones[j] >= i - j 可以从石头 j 走到石头 i
			if dp[j] && j+stones[j] >= i {
				dp[i] = true
				break
			}
		}
	}
	return dp[nums-1]
}
