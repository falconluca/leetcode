package leetcode

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	table := map[byte]int{}
	result, left := 0, 0
	for i := 0; i < len(s); i++ {
		cur := s[i]
		if _, ok := table[cur]; ok {
			// 出现重复字符时，指针指向重复字符下一个位置
			// 举例：abcb -> cb
			left = max(left, table[cur]+1)
		}
		table[cur] = i
		// i-left+1 为滑动窗口的长度
		result = max(result, i-left+1)
	}
	return result
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
