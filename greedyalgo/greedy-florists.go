package greedyalgo

import "sort"

// https://www.hackerrank.com/challenges/greedy-florist/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=greedy-algorithms
func getMinimumCost(k int32, c []int32) int32 {
	var i, j, n, noOfBuy, sum, temp int32
	n = int32(len(c))

	//sort temp array, complexity is O(n*log(n))
	sort.SliceStable(c, func(i, j int) bool {
		return c[i] < c[j]
	})

	for i = n - 1; i >= 0; i -= k {
		temp = i - k + 1
		if temp < 0 {
			temp = 0
		}
		for j = i; j >= temp; j-- {
			sum += (noOfBuy + 1) * c[j]
		}
		noOfBuy++
	}
	return sum
}
