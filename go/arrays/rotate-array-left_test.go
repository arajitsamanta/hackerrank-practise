package arrays

import (
	"reflect"
	"testing"
)

func TestGcd(t *testing.T) {
	a := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	b := []int32{4, 5, 6, 7, 8, 9, 10, 11, 12, 1, 2, 3}

	cases := []struct {
		in   []int32
		want []int32
	}{
		{a, b},
	}
	for _, c := range cases {
		got := rotateLeft(c.in, 3)
		if !reflect.DeepEqual(a, b) {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
