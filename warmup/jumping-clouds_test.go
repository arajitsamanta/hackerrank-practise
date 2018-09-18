package warmup

import (
	"testing"
)

func TestJumpingOnClouds(t *testing.T) {
	cases := []struct {
		in   []int32
		want int32
	}{
		{[]int32{0, 1, 0, 0, 0, 1, 0}, 3},
		{[]int32{0, 0, 1, 0, 0, 1, 0}, 4},
		{[]int32{0, 0, 0, 0, 1, 0}, 3},
	}
	for _, c := range cases {
		got := jumpingOnClouds(c.in)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
