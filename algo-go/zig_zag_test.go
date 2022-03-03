package leetcode_test

import (
	"leetcode"
	"testing"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		s       string
		numRows int
		result  string
	}{
		{
			"AB",
			1,
			"AB",
		},
		{
			"PAYPALISHIRING",
			3,
			"PAHNAPLSIIGYIR",
		},
		{
			"PAYPALISHIRING",
			4,
			"PINALSIGYAHRPI",
		},
		{
			"A",
			1,
			"A",
		},
	}

	for _, tt := range tests {
		r := leetcode.Convert(tt.s, tt.numRows)
		if r != tt.result {
			t.Fatalf("input %v, expect %v, but got %v", tt.s, tt.result, r)
		}
	}
}
