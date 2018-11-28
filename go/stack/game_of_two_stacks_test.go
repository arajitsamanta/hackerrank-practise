package stack

import "testing"

func TestGameOfTwoStacks(t *testing.T) {
	cases := []struct {
		x    int32
		in1  []int32
		in2  []int32
		want int32
	}{
		{10, []int32{4, 2, 4, 6, 1}, []int32{2, 1, 8, 5}, 4},
	}
	for _, c := range cases {
		got := twoStacks(c.x, c.in1, c.in2)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
