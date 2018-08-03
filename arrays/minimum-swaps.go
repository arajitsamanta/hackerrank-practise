package arrays

import (
	"fmt"
)

//https://www.hackerrank.com/challenges/minimum-swaps-2/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
func minimumSwaps(arr []int32) int32 {

	var i, length, swap int32
	length = int32(len(arr))
	fmt.Println(arr)
	//{4,3,1,2}
	for i = 0; i < length; i++ {

		if arr[i] == i+1 {
			continue
		}

		temp := arr[arr[i]-1]
		arr[arr[i]-1] = arr[i]
		arr[i] = temp
		//i = i - 1
		swap = swap + 1

		fmt.Printf("\npass %d val %v", i, arr)
	}
	return swap
}
