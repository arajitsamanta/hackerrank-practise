package greedyalgo

import (
	"sort"
)

//https://www.hackerrank.com/challenges/angry-children/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=greedy-algorithms
func maxMin(k int32, arr []int32) int32 {
	const (
		MaxInt32 = 1<<31 - 1
	)

	var (
		i, n, diff int32
		//temp       []int32
		minimum int32 = MaxInt32
	)

	//Sort the array first,  O(n*log(n))
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	n = int32(len(arr))

	//iterate over the sorted array upto n-k and check diff(max-min) of each arr of size k
	//O(n)
	for i = 0; i < n-k; i++ {
		diff = arr[i+k-1] - arr[i]
		if diff < minimum {
			minimum = diff
		}
	}

	return minimum
}
