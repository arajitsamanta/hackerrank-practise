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
		temp       []int32
		minimum    int32 = MaxInt32
	)

	//Sort the array first,  O(n*log(n))
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	n = int32(len(arr))

	//iterate over the sorted array and  slice it up to size k and check min/max of each slice
	//O(n)
	for i = 0; i < n; i++ {
		//check dif(max-min) only when slize size is k
		if n-i >= k {
			temp = arr[i : i+k]
			diff = temp[k-1] - temp[0]
		}

		if diff < minimum {
			minimum = diff
		}
	}

	return minimum
}
