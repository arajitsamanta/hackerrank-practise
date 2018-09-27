package maps

import "testing"

func TestTwoStrings(t *testing.T) {

	cases := []struct {
		in1  string
		in2  string
		want string
	}{
		{"a", "art", "YES"},
		{"be", "cat", "NO"},
		{"hello", "world", "YES"},
		{"hi", "world", "NO"},
	}
	for _, c := range cases {
		got := TwoString(c.in1, c.in2)
		if got != c.want {
			t.Errorf("Got %s, want %s", got, c.want)
		}
	}
}
