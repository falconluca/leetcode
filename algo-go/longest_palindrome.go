package leetcode

// LongestPalindrome 暴力破解
func LongestPalindrome(s string) string {
	if len(s) < 2 || len(s) > 1000 {
		return s
	}

	// 初始化最长回文串的起始位置和长度
	start, length := 0, 1

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if j-i+1 > length && ValidPalindrome(s, i, j) {
				// 发现新的最长回文子串啦
				length = j - i + 1
				start = i
			}
		}
	}

	return s[start : start+length]
}

// ValidPalindrome 判断 s[left...right] 是否为回文串
// 回文串是从左向右读，从右向左读都一样的串
// aaa, aba 是回文串
// bd 不是回文串
func ValidPalindrome(s string, left int, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
