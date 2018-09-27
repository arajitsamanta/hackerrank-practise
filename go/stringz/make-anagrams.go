package stringz

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

//https://www.hackerrank.com/challenges/ctci-making-anagrams/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=strings
func makeAnagrams(a string, b string) int32 {
	var count int32
	freq := make(map[string]int32)
	count = 0

	for _, char := range a {
		freq[string(char)]++
	}

	for _, char := range b {
		freq[string(char)]--
	}

	for _, v := range freq {
		count += abs(v)
	}

	return count
}
