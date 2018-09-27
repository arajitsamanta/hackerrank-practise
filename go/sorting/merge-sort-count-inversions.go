package sorting

import (
	"fmt"
)

// Left run is A[iLeft :iRight-1].
// Right run is A[iRight:iEnd-1  ].
func bottomUpMerge(A []int32, iLeft int32, iRight int32, iEnd int32, B []int32) int64 {
	var count int64
	i := iLeft
	j := iRight
	// While there are elements in the left or right runs...
	for k := iLeft; k < iEnd; k++ {
		// If left run head exists and is <= existing right run head.
		if i < iRight && (j >= iEnd || A[i] <= A[j]) {
			B[k] = A[i]
			i = i + 1
		} else {
			B[k] = A[j]
			j = j + 1
			count++
		}
	}
	return count
}

// Min returns the smaller of x or y.
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func MergeSortIterative(arr []int32) int64 {
	var i, width, n int32
	var inversions int64
	n = int32(len(arr))
	temp := make([]int32, n)

	// Each 1-element run in A is already "sorted".
	// Make successively longer sorted runs of length 2, 4, 8, 16... until whole array is sorted.
	for width = 1; width < n; width *= 2 {
		fmt.Printf("\nwidth:%d", width)
		for i = 0; i < n; i += 2 * width {
			fmt.Printf("\ni:%d", i)
			// Merge two runs: A[i:i+width-1] and A[i+width:i+2*width-1] to B[]
			// or copy A[i:n-1] to B[] ( if(i+width >= n) )
			//inversions = inversions + bottomUpMerge(arr, i, min(i+width, n), min(i+2*width, n), temp)
		}
		fmt.Println(temp)
		//fmt.Println(arr)
		copy(arr, temp)
	}

	fmt.Printf("sorted:%v", temp)
	return inversions

}

func merge(arr []int32, l int, m int, r int) int64 {
	fmt.Printf("\narr:%v l:%d m:%d r:%d", arr, l, m, r)
	//var temp int32
	var count int64
	var i, j, k int

	len1 := m - l + 1
	len2 := r - m

	tempA := make([]int32, len1)
	tempB := make([]int32, len2)

	for i = 0; i < len1; i++ {
		tempA[i] = arr[l+i]
	}
	for j = 0; j < len2; j++ {
		tempB[j] = arr[m+1+j]
	}
	fmt.Printf("\na:b -> %v %v", tempA, tempB)

	i = 0
	j = 0
	k = l
	for i < len1 && j < len2 {
		if tempA[i] <= tempB[j] {
			arr[k] = tempA[i]
			i++
		} else {
			arr[k] = tempB[j]
			j++
			count++
		}
		k++
	}
	/*
		for i := 0; i < lenA; i++ {
			for j := 0; j < lenB; j++ {
				if a[i] > b[j] {
					temp = a[i]
					a[i] = b[j]
					b[j] = temp
					count++
				}
			}
		}*/

	/* Copy the remaining elements of Left[], if there are any */
	for i < len1 {
		arr[k] = tempA[i]
		i++
		k++
	}

	/* Copy the remaining elements of Right[], if there are any */
	for j < len2 {
		arr[k] = tempB[j]
		j++
		k++
	}
	fmt.Printf("\nEach itr arr:%v,count %d", arr, count)

	return count
}

func CountInversions(arr []int32) int64 {
	var left, right, mid int
	var count int64
	n := len(arr)
	for width := 1; width < n-1; width *= 2 {

		// Pick starting point of different subarrays of current size
		for left = 0; left < n-1; left += 2 * width {
			// Find ending point of left subarray. mid+1 is starting
			// point of right
			mid = left + width - 1

			right = min(left+2*width-1, n-1)

			// Merge Subarrays arr[left_start...mid] & arr[mid+1...right_end]
			count += merge(arr, left, mid, right)
		}

		/*
			if width == 1 {
				for i = 0; i < n-1; i += 2 {
					if arr[i] > arr[i+1] {
						temp = arr[i]
						arr[i] = arr[i+1]
						arr[i+1] = temp
						count++
					}
				}
				fmt.Printf("case 1:%v, count:%d", arr, count)
			} else {
				start = 0
				end = 0
				length := (width * 2)
				for i = 0; i < n; i += length {
					start = i
					end = min(i+width, n)
					tempArr1 := arr[start:end]

					//second slice
					start = end
					end = min(end+width, n)
					tempArr2 := arr[start:end]
					fmt.Printf("\ntemp1:%v,temp2:%v", tempArr1, tempArr2)
					count += merge(tempArr1, tempArr2)
				}
			}*/

	}
	fmt.Printf("\nsorted array:%v", arr)

	return count
}
