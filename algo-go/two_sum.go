package leetcode

// 1. Two Sum
// https://leetcode.com/problems/two-sum/

func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for index, num := range nums {
		if i, ok := numMap[target-num]; ok {
			return []int{i, index}
		}
		numMap[num] = index
	}
	return nil
}
