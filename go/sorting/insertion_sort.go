package sorting

import "fmt"

// https://www.hackerrank.com/challenges/insertionsort2/problem?utm_campaign=challenge-recommendation&utm_medium=email&utm_source=24-hour-campaign
/*
Runtime Complexity
==================
Time: O(n^2)
Space:O(1)
*/
func insertionSort(n int32, arr []int32) []int32 {
	var i, j, key int32

	for i = 1; i < n; i++ {
		//Get the key to compare
		key = arr[i]
		j = i - 1

		//Insert key in the correct position
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
		printArray(n, arr)
	}

	return arr
}

func printArray(n int32, a []int32) {
	var i int32
	for i = 0; i < n; i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()
}
