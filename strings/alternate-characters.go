package strings

//https://www.hackerrank.com/challenges/alternating-characters/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=strings
func alternateCharacters(str string) int32 {
	var i, j, count int32
	n := int32(len(str))
	count = 0
	for i = 0; i < n; i++ {
		for j = i + 1; j < n; j++ {
			if j < n && str[i] == str[j] {
				count++
			} else {
				break
			}
		}
		i = j - 1
	}
	return count
}
