package leetcode_test

import (
	"leetcode"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			[]int{2, 3, -2, 4},
			6,
		},
		{
			[]int{-2, 0, -1},
			0,
		},
	}

	for _, tt := range tests {
		r := leetcode.MaxProduct(tt.in)
		if r != tt.out {
			t.Fatalf("input %v, execpt %v, but got %v\n", tt.in, tt.out, r)
		}
	}
}
