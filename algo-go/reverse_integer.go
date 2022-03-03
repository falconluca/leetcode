package leetcode

func Reverse(x int) int {
	result := 0
	for x != 0 {
		// 将整数当成一个队列，取低位
		last := x % 10
		// 12345 -> 1234
		x = x / 10

		// 处理溢出 -2^31 (-2147483648) <= x <= 2^31 - 1 (2147483647)
		// x <= 2147483647
		if result > 214748364 || (result == 214748364 && last > 7) {
			return 0
		}
		// x >= -2147483648
		if result < -214748364 || (result == -214748364 && last < -8) {
			return 0
		}

		// 新整型里低位转高位
		result = result*10 + last
	}
	return result
}
