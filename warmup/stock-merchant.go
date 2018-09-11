package warmup

//https://www.hackerrank.com/challenges/sock-merchant/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup
func sockMerchant(n int32, ar []int32) int32 {
	var i, result int32
	freq := make(map[int32]int32)
	for i = 0; i < n; i++ {
		freq[ar[i]]++
	}

	for _, v := range freq {
		result += v / 2
	}
	return result

}
