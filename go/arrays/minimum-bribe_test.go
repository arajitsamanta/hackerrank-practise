package arrays

import (
	"testing"
)

func TestMinimumBribe(t *testing.T) {

	a := []int32{1, 2, 5, 3, 7, 8, 6, 4}

	cases := []struct {
		in   []int32
		want int32
	}{
		{a, 3},
	}
	for _, c := range cases {
		MinimumBribe(c.in)
		/*if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}*/
	}
}
