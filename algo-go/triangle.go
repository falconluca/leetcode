package leetcode

// MinimumTotal 动态规划从三角形底下向上算
func MinimumTotal(triangle [][]int) int {
	// 从地下到上最小值之和的列表 状态压缩
	mini := make([]int, len(triangle)+1)

	// 从倒数第二行开始，逐层扫描三角形
	for i := len(triangle) - 1; i >= 0; i-- {
		// 逐列扫描三角形的每一层
		for j := 0; j < len(triangle[i]); j++ {
			// 当前行第j列的最小值 = 当前节点值 + 前面左右节点的最小值
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
