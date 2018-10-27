package dynamicprog

import (
	"testing"
)

func TestArrayMaxSubsetSum(t *testing.T) {
	cases := []struct {
		in   []int32
		want int32
	}{
		{[]int32{1, 2, 3, 4, 5}, 9},
		{[]int32{-3, 5, 2, 0, 4, -6}, 9},
		{[]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 30},
		{[]int32{-2, 1, 3, -4, 5}, 8},
	}
	for _, c := range cases {
		got := maxSubsetSum(c.in)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
