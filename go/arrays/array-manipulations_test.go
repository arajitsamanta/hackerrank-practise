package arrays

import (
	"testing"
)

func TestArrayManipulation(t *testing.T) {

	a := [][]int32{
		{1, 2, 100},
		{2, 5, 100},
		{3, 4, 100},
	}

	cases := []struct {
		in   [][]int32
		want int64
	}{
		{a, 200},
	}
	for _, c := range cases {
		got := ArrayManipulation(5, c.in)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
