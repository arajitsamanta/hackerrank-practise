package search

import (
	"testing"
)

func TestFindPossiblePairs(t *testing.T) {
	cases := []struct {
		arr  []int32
		k    int32
		want int32
	}{
		{[]int32{1, 2, 3, 4}, 1, 3},
		{[]int32{1, 5, 3, 4, 2}, 2, 3},
	}
	for _, c := range cases {
		got := pairs(c.k, c.arr)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
