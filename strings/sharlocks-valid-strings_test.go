package strings

import (
	"testing"
)

func TestSharlocksValidString(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"abcc", "YES"},
		{"abccc", "NO"},
		{"aabbcd", "NO"},
		{"aabbccddeefghi", "NO"},
		{"abcdefghhgfedecba", "YES"},
		{"aabbc", "YES"},
	}
	for _, c := range cases {
		got := isValid(c.in)
		if got != c.want {
			t.Errorf("Got %s, want %s", got, c.want)
		}
	}
}
