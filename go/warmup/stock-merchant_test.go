package warmup

import (
	"testing"
)

func TestStockMerchantCount(t *testing.T) {
	cases := []struct {
		len  int32
		arr  []int32
		want int32
	}{
		{7, []int32{1, 2, 1, 2, 1, 3, 2}, 2},
		{9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}, 3},
	}
	for _, c := range cases {
		got := sockMerchant(c.len, c.arr)
		if got != c.want {
			t.Errorf("Got %d, want %d", got, c.want)
		}
	}
}
