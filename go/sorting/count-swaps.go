package sorting

import (
	"fmt"
)

// https://www.hackerrank.com/challenges/ctci-bubble-sort/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=sorting
func CountSwaps(a []int32) {
	var temp, count int32
	n := len(a)

	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if a[j] > a[j+1] {
				temp = a[j]
				a[j] = a[j+1]
				a[j+1] = temp
				count++
			}
		}
	}

	fmt.Printf("Array is sorted in %d swaps.", count)
	fmt.Printf("\nFirst Element: %d", a[0])
	fmt.Printf("\nLast Element: %d", a[n-1])
}
