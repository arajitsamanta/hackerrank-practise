package search

import (
	"fmt"
	"sort"
)

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

func tripletsSmart(a []int32, b []int32, c []int32) int64 {
	var (
		count                        int64
		idxA, idxB, idxC, aDup, cDup int
		q                            int32
	)

	//Sort the arrays
	sortIt(a)
	sortIt(b)
	sortIt(c)

	lenA := len(a)
	lenB := len(b)
	lenC := len(c)

	//complexity is O(n^3)
	for idxB = 0; idxB < lenB; idxB++ {
		q = b[idxB]
		for idxA < lenA && a[idxA] <= q {
			// count duplicates in A array starting at index 1
			if idxA > 0 && a[idxA] == a[idxA-1] {
				aDup++
			}
			idxA++
		}

		for idxC < lenC && c[idxC] <= q {
			// count duplicates in C array starting at index 1
			if idxC > 0 && c[idxC] == c[idxC-1] {
				cDup++
			}
			idxC++
		}

		// Ignore duplicates in 'b' array; make sure terms are int64!
		if idxB == 0 || b[idxB-1] != b[idxB] {
			count += int64((idxA - aDup)) * int64((idxC - cDup))
		}

	}

	return count
}
