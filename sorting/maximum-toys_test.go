package sorting

import (
	"testing"
)

func TestMaximumToys(t *testing.T) {
	cases := []struct {
		in1  []int32
		in2  int32
		want int32
	}{
		{[]int32{1, 12, 5, 111, 200, 1000, 10}, 50, 4},
		{[]int32{1, 2, 3, 4}, 7, 3},
	}
	for _, c := range cases {
		got := MaximumToys(c.in1, c.in2)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
