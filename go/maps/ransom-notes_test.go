package maps

import "testing"

func TestCheckMagazine(t *testing.T) {

	cases := []struct {
		in1  []string
		in2  []string
		want string
	}{
		{
			[]string{"ive", "got", "a", "got", "lovely", "bunch", "of", "coconuts"}, []string{"ive", "got", "some", "coconuts"}, "No",
		},
		{
			[]string{"give", "me", "one", "grand", "today", "night"}, []string{"give", "one", "grand", "today"}, "Yes",
		},
		{
			[]string{"two", "times", "three", "is", "not", "four"}, []string{"two", "times", "is", "four"}, "Yes",
		},
	}
	for _, c := range cases {
		got := CheckMagazine(c.in1, c.in2)
		if got != c.want {
			t.Errorf("Got %s, want %s", got, c.want)
		}
	}
}
