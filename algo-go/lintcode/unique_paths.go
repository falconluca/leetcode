package lintcode

// https://www.lintcode.com/problem/114/
//
// A robot is located at the top-left corner of a `M x N` grid.
//
// The robot can only move either down or right at any point in time.
// The robot is trying to reach the bottom-right corner of the grid.
//
// How many possible unique paths are there?

func UniquePaths(m int, n int) int {
	// initial 2d array
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
