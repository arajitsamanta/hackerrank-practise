package warmup

import (
	"testing"
)

func TestRepeatedStrings(t *testing.T) {
	cases := []struct {
		in     string
		length int64
		want   int64
	}{
		{"aba", 10, 7},
		{"a", 1000000000000, 1000000000000},
	}
	for _, c := range cases {
		got := repeatedString(c.in, c.length)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
