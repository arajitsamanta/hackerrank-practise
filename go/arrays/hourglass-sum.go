package arrays

import "sort"

// hourglassSum - https://www.hackerrank.com/challenges/2d-array/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
func hourglassSum(arr [][]int32) int32 {
	var sum int32

	//iterate halfway for row and column
	rowMid := (6 / 2) + 1
	colMid := (6 / 2) + 1

	sumArr := make([]int, 16)

	sumArrIndex := 0

	for row := 0; row < rowMid; row++ {
		for col := 0; col < colMid; col++ {

			//sum of hourglass
			sum = sum + arr[row][col] + arr[row][col+1] + arr[row][col+2] + arr[row+1][col+1] + arr[row+2][col] + arr[row+2][col+1] + arr[row+2][col+2]

			//assign the sum of each hourglass to sum array
			sumArr[sumArrIndex] = int(sum)

			//increase sum array index
			sumArrIndex++

			//reset sum to 0 for next hourglass
			sum = 0
		}
	}

	//sor the array in natural order
	sort.Ints(sumArr)

	//retrun last element from the sum array
	return int32(sumArr[len(sumArr)-1])
}
