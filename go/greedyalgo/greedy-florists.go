package greedyalgo

import (
	"sort"
)

// https://www.hackerrank.com/challenges/greedy-florist/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=greedy-algorithms
func getMinimumCost(k int32, c []int32) int32 {
	var i, n, noOfBuy, sum int32
	n = int32(len(c))

	//We  need to optimize the order in which we purchase these flowers. The amount of additional money we need to pay later
	//is linear in . We want to buy the most expensive flowers first, at the lower multiple.
	//Sort in decending order
	sort.SliceStable(c, func(i, j int) bool {
		return c[j] < c[i]
	})

	for i = 0; i < n; i++ {
		sum += c[i] * (noOfBuy/k + 1)
		noOfBuy++
	}
	return sum
}
