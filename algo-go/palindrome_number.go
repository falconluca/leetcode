package leetcode

import "fmt"

func IsPalindrome(x int) bool {
	s := fmt.Sprint(x)
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
