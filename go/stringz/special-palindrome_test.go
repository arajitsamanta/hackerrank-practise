package stringz

import (
	"testing"
)

func TestSpecialPalindromeCount(t *testing.T) {
	cases := []struct {
		len  int32
		in   string
		want int64
	}{
		{8, "mnonopoo", 12},
		{5, "asasd", 7},
		{7, "abcbaba", 10},
		{4, "aaaa", 10},
	}
	for _, c := range cases {
		got := SubstrCount(c.len, c.in)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
