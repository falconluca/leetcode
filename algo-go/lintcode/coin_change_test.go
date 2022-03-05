package lintcode_test

import (
	"leetcode/lintcode"
	"testing"
)

func TestCoinChange(t *testing.T) {
	tests := []struct {
		coins  []int
		amount int
		output int
	}{
		{
			[]int{1, 2, 5},
			11,
			3,
		},
		{
			[]int{2},
			3,
			-1,
		},
	}

	for _, tt := range tests {
		r := lintcode.CoinChange(tt.coins, tt.amount)
		if r != tt.output {
			t.Fatalf("input coins: %v amount: %v, execpt %v, but got %v", tt.coins, tt.amount, tt.output, r)
		}
	}
}
