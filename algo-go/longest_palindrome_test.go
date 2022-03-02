package leetcode_test

import (
	"leetcode"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"babad",
			"bab",
		},
		{
			"cbbd",
			"bb",
		},
		{
			"ac",
			"a",
		},
		{
			"cbcbca",
			"cbcbc",
		},
	}

	for _, tt := range tests {
		r := leetcode.LongestPalindrome(tt.in)
		if r != tt.out {
			t.Fatalf("input %v, expect %v, but got %v", tt.in, tt.out, r)
		}
	}
}

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		s      string
		left   int
		right  int
		result bool
	}{
		{
			"aaa",
			0,
			2,
			true,
		},
		{
			"aba",
			0,
			2,
			true,
		},
		{
			"bd",
			0,
			1,
			false,
		},
	}

	for _, tt := range tests {
		r := leetcode.ValidPalindrome(tt.s, tt.left, tt.right)
		if r != tt.result {
			t.Fatalf("input %v, expect %v, but got %v", tt.s[tt.left:tt.right], tt.result, r)
		}
	}
}
