package leetcode_test

import (
	"leetcode"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"abcabcbb",
			3,
		},
		{
			"bbbbb",
			1,
		},
		{
			"pwwkew",
			3,
		},
	}

	for _, tt := range tests {
		r := leetcode.LengthOfLongestSubstring(tt.in)
		if r != tt.out {
			t.Fatalf("input %v, expect %v, but got %v", tt.in, tt.out, r)
		}
	}
}
