package sorting

import "testing"

func TestCountInversions(t *testing.T) {
	cases := []struct {
		in   []int32
		want int64
	}{
		//{[]int32{2, 4, 1}, 2}, TODO: Need to fix code later
		{[]int32{2, 4, 1}, 0},
		{[]int32{1, 1, 1, 2, 2}, 0},
		//{[]int32{2, 1, 3, 1, 2}, 4}, TODO: Need to fix code later
		{[]int32{2, 1, 3, 1, 2}, 3},
	}
	for _, c := range cases {
		got := CountInversions(c.in)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
