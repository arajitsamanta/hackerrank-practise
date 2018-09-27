package search

import (
	"testing"
)

func TestFindPossibleTriplets(t *testing.T) {
	cases := []struct {
		a    []int32
		b    []int32
		c    []int32
		want int64
	}{
		{[]int32{1, 3, 5}, []int32{2, 3}, []int32{1, 2, 3}, 8},
		{[]int32{3, 5, 7}, []int32{3, 6}, []int32{4, 6, 9}, 4},
		{[]int32{5, 4, 1}, []int32{2, 3, 3}, []int32{3, 2, 1}, 5},
		{[]int32{1, 4, 5, 7}, []int32{5, 7, 9}, []int32{7, 9, 11, 13}, 12},
	}
	for _, c := range cases {
		got := tripletsSmart(c.a, c.b, c.c)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
