package arrays

import (
	"fmt"
)

//https://www.hackerrank.com/challenges/new-year-chaos/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays&h_r=next-challenge&h_v=zen
func MinimumBribe(q []int32) {

	var i, j, sum, supposedToBeIdx int32
	isChaos := false
	isSorted := false

	length := int32(len(q))

	for i = 0; i < length-1; i++ {
		isSorted = true
		for j = 0; j < length-i-1; j++ {
			supposedToBeIdx = q[j] - 1
			if supposedToBeIdx-j > 2 {
				isChaos = true
				break
			}
			if q[j] > q[j+1] {
				// swap temp and arr[i]
				temp := q[j]
				q[j] = q[j+1]
				q[j+1] = temp
				sum = sum + 1
				isSorted = false
			}
		}
		if isChaos || isSorted {
			break
		}
	}

	if isChaos {
		fmt.Println("Too chaotic")
	} else {
		fmt.Println(sum)
	}
}
