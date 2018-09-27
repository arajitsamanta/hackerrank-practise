package greedyalgo

import (
	"testing"
)

func TestMaximumLuckBalance(t *testing.T) {
	cases := []struct {
		in   [][]int32
		k    int32
		want int32
	}{
		{
			[][]int32{
				{5, 1},
				{1, 1},
				{4, 0},
			}, 2, 10,
		},
		{
			[][]int32{
				{5, 1},
				{1, 1},
				{4, 0},
			}, 1, 8,
		},
		{
			[][]int32{
				{5, 1},
				{2, 1},
				{1, 1},
				{8, 1},
				{10, 0},
				{5, 0},
			}, 3, 29,
		},
	}
	for _, c := range cases {
		got := luckBalance(c.k, c.in)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
