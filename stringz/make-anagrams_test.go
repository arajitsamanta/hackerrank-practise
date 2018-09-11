package stringz

import (
	"testing"
)

func TestMakeAnagrams(t *testing.T) {
	cases := []struct {
		in1  string
		in2  string
		want int32
	}{
		{"abc", "bcd", 2},
		{"cde", "abc", 4},
	}
	for _, c := range cases {
		got := makeAnagrams(c.in1, c.in2)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
