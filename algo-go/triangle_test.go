package leetcode_test

import (
	"leetcode"
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	tests := []struct {
		in  [][]int
		out int
	}{
		{
			[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
			11,
		},
		{
			[][]int{{-10}},
			-10,
		},
	}

	for _, tt := range tests {
		r := leetcode.MinimumTotal4(tt.in)
		if r != tt.out {
			t.Fatalf("input %v, execpt %v, but got %v\n", tt.in, tt.out, r)
		}
	}
}
