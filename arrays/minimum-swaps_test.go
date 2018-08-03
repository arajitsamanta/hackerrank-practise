package arrays

import (
	"testing"
)

func TestMinimumSwap(t *testing.T) {
	//a := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	b := []int32{4, 3, 1, 2}

	cases := []struct {
		in   []int32
		want int32
	}{
		{b, 3},
	}
	for _, c := range cases {
		got := minimumSwaps(c.in)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
