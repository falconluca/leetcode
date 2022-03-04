package leetcode

// MinimumTotal 动态规划从三角形底下向上算
func MinimumTotal(triangle [][]int) int {
	// 从地下到上最小值之和的列表 状态压缩
	mini := make([]int, len(triangle)+1)

	for i := len(triangle) - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			mini[j] = triangle[i][j] + min(mini[j], mini[j+1])
		}
	}
	return mini[0]
}

func min(prev int, cur int) int {
	if prev > cur {
		return cur
	}
	return prev
}
