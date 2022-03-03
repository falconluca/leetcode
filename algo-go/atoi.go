package leetcode

import (
	"math"
	"unicode"
)

// Atoi
// 0 <= s.length <= 200
// 32 位有符号整数
func Atoi(s string) int {
	if s == "" {
		return 0
	}

	result := 0
	ptr := 0

	// 处理空格
	for ptr < len(s) && unicode.IsSpace(rune(s[ptr])) {
		ptr++
	}

	// 处理符号
	ne := false
	if ptr < len(s) && (s[ptr] == '-' || s[ptr] == '+') {
		ne = s[ptr] == '-'
		ptr++
	}

	var digit int
	for ptr < len(s) && unicode.IsDigit(rune(s[ptr])) {
		digit = int(s[ptr] - '0')

		// 处理溢出
		if ne {
			//  -2^31 (-2147483648) <= x <= 2^31 - 1 (2147483647)
			if -result < (math.MinInt32+digit)/10 {
				return math.MinInt32
			}
		} else {
			if result > (math.MaxInt32-digit)/10 {
				return math.MaxInt32
			}
		}

		result = 10*result + digit
		ptr++
	}

	if ne {
		return -result
	}
	return result
}
