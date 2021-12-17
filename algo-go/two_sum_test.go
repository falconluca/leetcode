package leetcode_test

import (
	"leetcode"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct{
		nums []int
		target int
		expect []int
	}{
		{
			nums: []int{2, 7, 11, 15},
			target: 9,
			expect: []int{0, 1},
		},
		{
			nums: []int{3, 2, 4},
			target: 6,
			expect: []int{1, 2},
		},
		{
			nums: []int{3, 3},
			target: 6,
			expect: []int{0, 1},
		},
	}

	for _, tt := range tests {
		actual := leetcode.TwoSum(tt.nums, tt.target)
		if !isSliceEquals(tt.expect, actual) {
			t.Fatalf("nums is %v, target is %v and expect %v, but got %v",
				tt.nums, tt.target, tt.expect, actual)
		}
	}
}

func isSliceEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}