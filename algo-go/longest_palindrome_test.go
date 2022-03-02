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
	}

	for _, tt := range tests {
		r := leetcode.LongestPalindrome(tt.in)
		if r != tt.out {
			t.Fatalf("input %v, expect %v, but got %v", tt.in, tt.out, r)
		}
	}
}

func TestValidPalindrome(t *testing.T) {
	// TODO
}
