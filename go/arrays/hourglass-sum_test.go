package arrays

import (
	"testing"
)

func TestHourglassSum(t *testing.T) {

	a := [][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}

	cases := []struct {
		in   [][]int32
		want int32
	}{
		{a, 18},
	}
	for _, c := range cases {
		got := hourglassSum(c.in)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
