package maps

import "testing"

func TestCountTriplets(t *testing.T) {

	cases := []struct {
		arr  []int64
		gp   int64
		want int64
	}{
		{
			[]int64{1, 2, 1, 2, 4}, 2, 3,
		},
		{
			[]int64{1, 3, 9, 9, 27, 81}, 3, 6,
		},
		{
			[]int64{1, 5, 5, 25, 125}, 5, 4,
		},
		{
			[]int64{1, 2, 2, 4}, 2, 2,
		},
		{
			[]int64{55, 55, 55, 55}, 1, 4,
		},
	}
	for _, c := range cases {
		got := CountTriplets(c.arr, c.gp)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
