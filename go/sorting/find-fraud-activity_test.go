package sorting

import (
	"testing"
)

func TestFraudActivityNotifications(t *testing.T) {
	cases := []struct {
		in1  []int32
		in2  int32
		want int32
	}{
		{[]int32{2, 3, 4, 2, 3, 6, 8, 4, 5}, 5, 2},
		{[]int32{10, 20, 30, 40, 50}, 3, 1},
		{[]int32{1, 2, 3, 4, 4}, 4, 0},
	}
	for _, c := range cases {
		got := ActivityNotifications(c.in1, c.in2)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
