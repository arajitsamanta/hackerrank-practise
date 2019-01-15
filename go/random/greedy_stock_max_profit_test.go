package random

import (
	"testing"
)

func TestStockMaxProfit(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{10, 7, 5, 8, 11, 9}, 6},
		{[]int{10, 5, 2, 1}, -1},
	}
	for _, c := range cases {
		got := getMaxProfit(c.in)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
