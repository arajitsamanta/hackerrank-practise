package sorting

import (
	"sort"
)

//https://www.hackerrank.com/challenges/mark-and-toys/forum?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=sorting
func MaximumToys(prices []int32, k int32) int32 {
	var count int32
	sort.SliceStable(prices, func(i, j int) bool {
		return prices[i] < prices[j]
	})

	n := len(prices)
	for i := 0; i < n; i++ {
		if prices[i] < k {
			k -= prices[i]
			count++
		}
	}
	return count
}
