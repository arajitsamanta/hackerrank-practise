package sorting

import (
	"fmt"
)

func print(n int32, arr []int32) {
	var i int32
	for i = 0; i < n; i++ {
		if i == n-1 {
			fmt.Printf("%d", arr[i])
		} else {
			fmt.Printf("%d ", arr[i])
		}
	}
	fmt.Println()
}

//https://www.hackerrank.com/challenges/insertionsort1/problem?utm_campaign=challenge-recommendation&utm_medium=email&utm_source=24-hour-campaign
func sortedInsertionSort(n int32, arr []int32) {
	var i, target int32
	target = arr[n-1]
	for i = n - 2; i >= 0; i-- {
		if arr[i] < target {
			arr[i+1] = target
			print(n, arr)
			break
		} else {
			arr[i+1] = arr[i]
			print(n, arr)
		}
	}

	if i == -1 {
		arr[0] = target
		print(n, arr)
	}

}
