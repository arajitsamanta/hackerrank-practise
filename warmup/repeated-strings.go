package warmup

// https://www.hackerrank.com/challenges/repeated-string/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup
func repeatedString(s string, n int64) int64 {
	var i, length, count, remain int64
	length = int64(len(s))

	//Get initial count of 'a' in origicnal string
	for _, ch := range s {
		if string(ch) == "a" {
			count++
		}
	}

	//Total count should be n/length of s * count
	count = (n / length) * count

	// check if there is any remainder and add it to the final count
	remain = n % length
	for i = 0; i < remain; i++ {
		if string(s[i]) == "a" {
			count++
		}
	}
	return count
}
