package warmup

//https://www.hackerrank.com/challenges/jumping-on-the-clouds/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup&h_r=next-challenge&h_v=zen
func jumpingOnClouds(c []int32) int32 {
	var count int32
	n := len(c)
	for i := 0; i < n-1; {
		i += 2
		if i < n && c[i] == 1 {
			i--
		}
		count++
	}
	return count
}
