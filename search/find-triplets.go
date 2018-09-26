package search

import (
	"fmt"
	"sort"
)

//Remove duplicate from an array
func removeDuplicate(arr []int32) []int32 {
	keys := make(map[int32]bool)
	list := []int32{}
	for _, entry := range arr {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

//https://www.hackerrank.com/challenges/triple-sum/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=search
func tripletsBruteForce(a []int32, b []int32, c []int32) int64 {
	var (
		count int64
	)

	a = removeDuplicate(a)
	b = removeDuplicate(b)
	c = removeDuplicate(c)

	lenA := len(a)
	lenB := len(b)
	lenC := len(c)

	//complexity is O(n^3)
	for i := 0; i < lenA; i++ {
		for j := 0; j < lenB; j++ {
			for k := 0; k < lenC; k++ {
				if a[i] <= b[j] && b[j] >= c[k] {
					fmt.Println(a[i], b[j], c[k])
					count++
				}
			}
		}
	}
	return count
}

func sortIt(arr []int32) {
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}

func binarySearch(arr []int32, n int, toSearch int32, flagLeftSearch bool) int {
	var (
		mid   int
		found bool
	)
	start := 0
	end := n - 1
	//n = 3

	for start <= end {

		mid = (start + end) / 2

		if arr[mid] == toSearch {
			found = true
			break
		}

		if arr[mid] > toSearch {
			end = mid - 1
		} else {
			start = mid + 1
		}

	}

	fmt.Println(mid, arr[mid])

	/*if flagLeftSearch {
		if toSearch <= arr[mid] {
			return (n - mid + 1)
		}
	} else {
		if toSearch >= arr[mid] {
			return (n - mid + 1)
		}
	}*/

	if found {
		return mid + 1
	}

	if flagLeftSearch {
		if toSearch <= arr[mid] {
			return (n - mid + 1)
		}
	} else {
		if toSearch >= arr[mid] {
			return (n - mid + 1)
		}
	}

	return 0
}

func tripletsSmart(a []int32, b []int32, c []int32) int64 {
	var (
		count int64
	)

	//Remove duplicates
	a = removeDuplicate(a)
	b = removeDuplicate(b)
	c = removeDuplicate(c)

	//Sort the arrays
	sortIt(a)
	sortIt(b)
	sortIt(c)
	fmt.Println(a, b, c)

	lenA := len(a)
	lenB := len(b)
	lenC := len(c)

	//complexity is O(n^3)
	for i := 0; i < lenB; i++ {
		q := b[i]

		lCount := binarySearch(a, lenA, q, true)
		fmt.Println("left:", lCount)

		rCount := binarySearch(c, lenC, q, false)
		fmt.Println("right", rCount)
		count += int64(lCount * rCount)

	}

	return count
}
