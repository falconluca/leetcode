package leetcode

var res int

func MaxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	result, curMax, curMin := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		curMax, curMin = curMax*nums[i], curMin*nums[i]
		curMin, curMax = minN(curMax, curMin, nums[i]), maxN(curMax, curMin, nums[i])
		if curMax > result {
			result = curMax
		}
	}
	return result
}

func minN(a int, b int, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	}

	if b < c {
		return b
	}
	return c
}

func maxN(a int, b int, c int) int {
	if a > b {
		if a > c {
			return a
		} else {
			return c
		}
	}

	if b > c {
		return b
	}
	return c
}

func recursion(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	res = nums[0]
	dfs(nums, 0, nums[0])
	return res
}

// dfs
func dfs(nums []int, i int, value int) {
	if i >= len(nums)-1 {
		res = max(res, value)
		return
	}
	res = max(res, value)
	dfs(nums, i+1, value*nums[i+1])
	dfs(nums, i+1, nums[i+1])
}
