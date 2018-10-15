package sorting

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	cases := []struct {
		n    int32
		in   []int32
		want []int32
	}{
		{
			6, []int32{5, 2, 4, 6, 1, 3}, []int32{1, 2, 3, 4, 5, 7},
		},
	}
	for _, c := range cases {
		got := insertionSort(c.n, c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Got %v, want %v", got, c.want)
		}
	}
}
