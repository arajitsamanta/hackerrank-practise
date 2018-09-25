package greedyalgo

import (
	"testing"
)

func TestMinimumfairArray(t *testing.T) {
	cases := []struct {
		in   []int32
		k    int32
		want int32
	}{
		{[]int32{1, 4, 7, 2}, 2, 1},
		{[]int32{10, 100, 300, 200, 1000, 20, 30}, 3, 20},
		{[]int32{1, 2, 3, 4, 10, 20, 30, 40, 100, 200}, 4, 3},
		{[]int32{1, 2, 1, 2, 1}, 2, 0},
	}
	for _, c := range cases {
		got := maxMin(c.k, c.in)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
