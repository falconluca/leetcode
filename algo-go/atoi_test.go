package leetcode_test

import (
	"leetcode"
	"testing"
)

func TestAtoi(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"21474836460",
			2147483647,
		},
		{
			"42",
			42,
		},
		{
			"-42",
			-42,
		},
		{
			"4193 with words",
			4193,
		},
	}

	for _, tt := range tests {
		r := leetcode.Atoi(tt.in)
		if r != tt.out {
			t.Fatalf("input %v, execpt %v, but got %v", tt.in, tt.out, r)
		}
	}
}
