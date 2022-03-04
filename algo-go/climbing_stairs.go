package leetcode

// ClimbStairs n 层楼梯，每次只能爬 1 或 2 个台阶，有多少种爬法
func ClimbStairs(n int) int {
	// 0层 0种爬法
	// 1层 1种爬法 1
	// 2层 2种爬法 2或1+1
	if n <= 2 {
		return n
	}

	opt := make([]int, n)
	// 1层 1种爬法 1
	opt[0] = 1
	// 2层 2种爬法 2或1+1
	opt[1] = 2

	for i := 2; i < n; i++ {
		// 第i层的爬法 = 第i-1层的爬法 + 第i-2层的爬法
		opt[i] = opt[i-1] + opt[i-2]
	}

	// 第n层的爬法 因为从0开始
	return opt[n-1]
}
