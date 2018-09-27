package arrays

//https://www.hackerrank.com/challenges/crush/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
func ArrayManipulation(n int32, queries [][]int32) int64 {
	var max, sum int64
	var j int32
	arr := make([]int64, n+1)

	//Prefix sum array algorithim
	for i := 0; i < len(queries); i++ {
		col := queries[i]
		left := col[0]
		right := col[1] + 1
		arr[left] += int64(col[2])
		if right <= n {
			arr[right] -= int64(col[2])
		}
	}

	for j = 1; j <= n; j++ {
		sum = sum + arr[j]
		if max < sum {
			max = sum
		}
	}

	return max
}
