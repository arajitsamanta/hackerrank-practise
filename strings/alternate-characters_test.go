package strings

import (
	"testing"
)

func TestAlternateCharacters(t *testing.T) {
	cases := []struct {
		in   string
		want int32
	}{
		{"AABAAB", 2},
		{"AAAA", 3},
		{"BBBBB", 4},
		{"ABABABAB", 0},
		{"BABABA", 0},
		{"AAABBB", 4},
	}
	for _, c := range cases {
		got := alternateCharacters(c.in)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
