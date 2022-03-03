package leetcode

import "fmt"

// Convert 表格 + 翻转指针
func Convert(s string, numRows int) string {
	if len(s) < 1 || len(s) > 1000 {
		return s
	}
	if numRows < 2 || numRows > 1000 {
		return s
	}

	// initial list
	list := []string{}
	for i := 0; i < numRows; i++ {
		list = append(list, "")
	}

	cur, pos := 0, -1
	for i := 0; i < len(s); i++ {
		ch := string(s[i])
		substr := list[cur]
		list[cur] = fmt.Sprint(substr, ch)

		// 指针走到头或尾的时候，翻转指针
		if cur == 0 || cur == numRows-1 {
			pos = -pos
		}

		cur += pos
	}

	result := ""
	for i := 0; i < numRows; i++ {
		result = fmt.Sprint(result, list[i])
	}
	return result
}
