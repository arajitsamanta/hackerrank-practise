package arrays

import (
	"testing"
)

func TestMinimumSwap(t *testing.T) {
	//a := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	//b := []int32{4, 3, 1, 2}

	//b := []int32{7, 1, 3, 2, 4, 5, 6}

	b := []int32{8, 3, 5, 2, 4, 1, 6}

	cases := []struct {
		in   []int32
		want int32
	}{
		{b, 5},
	}
	for _, c := range cases {
		got := MinimumSwaps(c.in)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
