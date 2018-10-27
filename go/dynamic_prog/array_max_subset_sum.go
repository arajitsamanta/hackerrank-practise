package dynamicprog

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

// Complete the maxSubsetSum function below.
func maxSubsetSum(arr []int32) int32 {

	var (
		i, n int32
	)

	n = int32(len(arr))

	if n == 0 {
		return 0
	}

	arr[0] = max(0, arr[0])

	if n == 1 {
		return arr[0]
	}
	arr[1] = max(arr[0], arr[1])

	for i = 2; i < n; i++ {
		arr[i] = max(arr[i-1], arr[i]+arr[i-2])
	}
	return arr[n-1]
}
