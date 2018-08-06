package arrays

import (
	"fmt"
)

//https://www.hackerrank.com/challenges/minimum-swaps-2/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
func MinimumSwaps(arr []int32) int32 {

	var i, length, count int32

	length = int32(len(arr))

	for i = 0; i < length; {

		fmt.Printf("\npass %d val %v", i, arr)

		if arr[i] == (i+1) || arr[i] >= length {
			i++
			continue
		}

		tmp := arr[i]
		arr[i] = arr[tmp-1]
		arr[tmp-1] = tmp

		count++
	}

	return count
}
