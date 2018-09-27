package maps

import (
	"reflect"
	"testing"
)

func TestFreqQuery(t *testing.T) {

	cases := []struct {
		in   [][]int32
		want []int32
	}{
		{
			[][]int32{
				{1, 1},
				{2, 2},
				{3, 2},
				{1, 1},
				{1, 1},
				{2, 1},
				{3, 2},
			}, []int32{0, 1},
		},
		{
			[][]int32{
				{1, 5},
				{1, 6},
				{3, 2},
				{1, 10},
				{1, 10},
				{1, 6},
				{2, 5},
				{3, 2},
			}, []int32{0, 1},
		},
		{
			[][]int32{
				{3, 4},
				{2, 1003},
				{1, 16},
				{3, 1},
			}, []int32{0, 1},
		},
		{
			[][]int32{
				{1, 3},
				{2, 3},
				{3, 2},
				{1, 4},
				{1, 5},
				{1, 5},
				{1, 4},
				{3, 2},
				{2, 4},
				{3, 2},
			}, []int32{0, 1, 1},
		},
	}

	for _, c := range cases {
		got := FreqQuery(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
