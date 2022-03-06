package lintcode_test

import (
	"leetcode/lintcode"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		n      int
		m      int
		output int
	}{
		{
			1,
			3,
			1,
		},
		{
			3,
			3,
			6,
		},
	}

	for _, tt := range tests {
		r := lintcode.UniquePaths(tt.n, tt.m)
		if r != tt.output {
			t.Fatalf("input m: %v n: %v, execpt %v, but got %v\n", tt.n, tt.m, tt.output, r)
		}
	}
}
