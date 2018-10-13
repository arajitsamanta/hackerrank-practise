package sorting

import "testing"

func TestSortedInsertionSort(t *testing.T) {
	cases := []struct {
		in []int32
		n  int32
	}{
		{[]int32{1, 2, 4, 5, 3}, 5},
		{[]int32{2, 4, 6, 8, 3}, 5},
		{[]int32{2, 3, 4, 5, 6, 7, 8, 9, 10, 1}, 10},
	}
	for _, c := range cases {
		sortedInsertionSort(c.n, c.in)
		t.Errorf("failed")
	}
}
