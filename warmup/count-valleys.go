package warmup

//https://www.hackerrank.com/challenges/counting-valleys/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup
func countingValleys(n int32, s string) int32 {
	var i, count, level int32
	count = 0
	for i = 0; i < n; i++ {
		if s[i] == 'U' {
			level++
		}

		if s[i] == 'D' {
			level--
		}

		//We only care about the number of valleys. So jwe just need ti figure out the number of times you came back up to sea level.
		if level == 0 && s[i] == 'U' {
			count++
		}
	}
	return count
}
