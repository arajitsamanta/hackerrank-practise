package greedyalgo

import (
	"testing"
)

func TestMinimumAbsoluteDifference(t *testing.T) {
	cases := []struct {
		in   []int32
		want int32
	}{
		{[]int32{3, 7, 0}, 3},
		{[]int32{-2, 2, 4}, 2},
		{[]int32{-59, -36, -13, 1, -53, -82, -2, -96, -54, 75}, 1},
		{[]int32{1, -3, 71, 68, 17}, 3},
	}
	for _, c := range cases {
		got := minimumAbsoluteDifference(c.in)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
