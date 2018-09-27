package search

import (
	"testing"
)

func TestIceCreamFlavour(t *testing.T) {
	cases := []struct {
		costs []int32
		money int32
		want  int32
	}{
		{[]int32{1, 2, 3, 4}, 3, 11},
		{[]int32{2, 5, 6}, 3, 13},
		{[]int32{2, 5, 6}, 2, 15},
		{[]int32{1, 3, 5, 7, 9}, 3, 29},
	}
	for _, c := range cases {
		whatFlavors(c.costs, c.money)
	}
}
