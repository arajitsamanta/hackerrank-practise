package maps

import "testing"

func TestFindAnagrams(t *testing.T) {

	cases := []struct {
		in   string
		want int32
	}{
		{"abba", 4},
		{"mom", 2},
		{"kkkk", 10},
		{"abcd", 0},
		{"ifailuhkqq", 3},
	}
	for _, c := range cases {
		got := FindAnagrams(c.in)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
