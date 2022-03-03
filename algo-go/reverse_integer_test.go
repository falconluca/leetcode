package leetcode_test

import (
	"fmt"
	"leetcode"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{
			123,
			321,
		},
		{
			-123,
			-321,
		},
		{
			120,
			21,
		},
		{
			0,
			0,
		},
	}

	for _, tt := range tests {
		r := leetcode.Reverse(tt.in)
		if r != tt.out {
			t.Fatalf("input %v, except %v, but got %v\n", tt.in, tt.out, r)
		}
	}
}

func ExampleReverse() {
	x := 12345
	x /= 10
	fmt.Println(x)
	// Output:
	// 1234
}
