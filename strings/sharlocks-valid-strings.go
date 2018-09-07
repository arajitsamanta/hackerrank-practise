package strings

//https://www.hackerrank.com/challenges/sherlock-and-valid-string/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=strings
func isValid(s string) string {
	var minDelete bool
	var result string
	freq := make(map[string]int32)

	for _, char := range s {
		//put char in frequency character map and value in value map
		freq[string(char)]++
	}

	// Convert map to slice of values.
	values := []int32{}
	for _, value := range freq {
		values = append(values, value)
	}

	first := values[0]
	result = "YES"
	for i, val := range values {
		if val != first && !minDelete {
			val--
			values[i] = val
			minDelete = true
		}

		if first != int32(val) && val != 0 {
			result = "NO"
			break
		}
	}

	return result
}
