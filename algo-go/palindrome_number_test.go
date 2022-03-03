package leetcode_test

import (
	"leetcode"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		in  int
		out bool
	}{
		{
			121,
			true,
		},
		{
			-121,
			false,
		},
		{
			10,
			false,
		},
	}

	for _, tt := range tests {
		r := leetcode.IsPalindrome(tt.in)
		if r != tt.out {
			t.Fatalf("input %v, execpt %v, but got %v\n", tt.in, tt.out, r)
		}
	}
}
