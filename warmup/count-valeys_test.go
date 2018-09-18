package warmup

import (
	"testing"
)

func TestCountingValleys(t *testing.T) {
	cases := []struct {
		n    int32
		in   string
		want int32
	}{
		{8, "UDDDUDUU", 1},
		{8, "DDUUUUDD", 1},
	}
	for _, c := range cases {
		got := countingValleys(c.n, c.in)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
