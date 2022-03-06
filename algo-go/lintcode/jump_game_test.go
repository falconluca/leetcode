package lintcode_test

import (
	"leetcode/lintcode"
	"testing"
)

func TestCanJump(t *testing.T) {
	tests := []struct {
		input  []int
		output bool
	}{
		{
			[]int{2, 3, 1, 1, 4},
			true,
		},
		{
			[]int{3, 2, 1, 0, 4},
			false,
		},
	}

	for _, tt := range tests {
		r := lintcode.CanJump(tt.input)
		if r != tt.output {
			t.Fatalf("input %v, execpt %v, but got %v\n", tt.input, tt.output, r)
		}
	}
}
