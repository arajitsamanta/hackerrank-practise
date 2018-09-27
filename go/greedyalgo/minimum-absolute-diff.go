package greedyalgo

import "sort"

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

// https://www.hackerrank.com/challenges/minimum-absolute-difference-in-an-array/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=greedy-algorithms
func minimumAbsoluteDifferenceBruteForce(arr []int32) int32 {
	var (
		min  int32 = 2147483647 // assign int32 max value as minimum
		temp int32
	)
	n := len(arr)
	//Complexity is O(n^2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] != arr[j] {
				temp = abs(arr[i] - arr[j])
				if temp < min {
					min = temp
				}
			}
		}
	}

	return min
}

func minimumAbsoluteDifference(arr []int32) int32 {
	var (
		min  int32 = 2147483647 // assign int32 max value as minimum
		temp int32
	)

	//sort original array, complexity is O(n*log(n))
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	n := len(arr)
	for i := 0; i < n; i++ {
		if (i + 1) < n {
			temp = abs(arr[i] - arr[i+1])
			if temp < min {
				min = temp
			}
		}
	}

	return min
}
