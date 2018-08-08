package maps

//https://www.hackerrank.com/challenges/ctci-ransom-note/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps
func CheckMagazine(magazine []string, note []string) string {

	words := make(map[string]int)

	for i := 0; i < len(magazine); i++ {
		key := magazine[i]
		words[key]++
	}

	length := len(note)

	var result string

	for j := 0; j < length; j++ {
		key := note[j]
		v, ok := words[key]
		if ok && 0 != v {
			words[key]--
		} else {
			result = "No"
			break
		}

		if j == length-1 {
			result = "Yes"
		}
	}

	return result

}
