package leetcode_test

import (
	"leetcode"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			[]int{7, 1, 5, 3, 6, 4},
			5,
		},
		{
			[]int{7, 6, 4, 3, 1},
			0,
		},
	}

	for _, tt := range tests {
		r := leetcode.MaxProfit(tt.in)
		if r != tt.out {
			t.Fatalf("input %v, expect %v, but got %v\n", tt.in, tt.out, r)
		}
	}
}
