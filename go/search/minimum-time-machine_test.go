package search

import (
	"testing"
)

func TestMinimumTimeRequired(t *testing.T) {
	cases := []struct {
		machines []int64
		goal     int64
		want     int64
	}{
		{[]int64{2, 3, 2}, 10, 8},
		{[]int64{2, 3}, 5, 6},
		{[]int64{1, 3, 4}, 10, 7},
		{[]int64{4, 5, 6}, 12, 20},
	}
	for _, c := range cases {
		got := MinTimeBruteForce(c.machines, c.goal)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
