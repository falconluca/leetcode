package leetcode_test

import (
	"leetcode"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{
			2,
			2,
		},
		{
			3,
			3,
		},
	}

	for _, tt := range tests {
		r := leetcode.ClimbStairs(tt.in)
		if r != tt.out {
			t.Fatalf("input %v, execpt %v, but got %v\n", tt.in, tt.out, r)
		}
	}
}
