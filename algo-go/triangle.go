package leetcode

import "math"

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

var minimum = math.MaxInt

// dfs：遍历
// 2^n -> 1 2 4 8
func MinimumTotal2(triangle [][]int) int {
	MinimumTotal_DFS_Traverse(triangle, 0, 0, 0)
	return minimum
}

func MinimumTotal_DFS_Traverse(triangle [][]int, x int, y int, pathSum int) {
	if len(triangle) == x { // 越界...
		minimum = min(pathSum, minimum)
		return
	}

	MinimumTotal_DFS_Traverse(triangle, x+1, y, pathSum+triangle[x][y])
	MinimumTotal_DFS_Traverse(triangle, x+1, y+1, pathSum+triangle[x][y])
}

// dfs：Divide Conquer
func MinimumTotal3(triangle [][]int) int {
	return MinimumTotal_DFS_DC(triangle, 0, 0)
}

func MinimumTotal_DFS_DC(triangle [][]int, x int, y int) int {
	if x == len(triangle) {
		return 0
	}

	left := MinimumTotal_DFS_DC(triangle, x+1, y)
	right := MinimumTotal_DFS_DC(triangle, x+1, y+1)
	return min(left, right) + triangle[x][y]
}

type MemoKey struct {
	x int
	y int
}

// dfs：Divide Conquer + memo
// 2^n -> n^2
// memoization search 记忆化搜索
func MinimumTotal4(triangle [][]int) int {
	len := len(triangle)
	memo := make(map[MemoKey]int, len*len/2)
	return MinimumTotal_DFS_DC_Memo(triangle, 0, 0, memo)
}

func MinimumTotal_DFS_DC_Memo(triangle [][]int, x int, y int, memo map[MemoKey]int) int {
	if x == len(triangle) {
		return 0
	}

	val, ok := memo[MemoKey{x, y}]
	if ok {
		return val
	}

	left := MinimumTotal_DFS_DC(triangle, x+1, y)
	right := MinimumTotal_DFS_DC(triangle, x+1, y+1)
	memo[MemoKey{x, y}] = min(left, right) + triangle[x][y]
	return memo[MemoKey{x, y}]
}
