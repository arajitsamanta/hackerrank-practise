package maps

import (
	"sort"
)

func stringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func sortStringByCharacter(s string) string {
	r := stringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func countAnagrams(words []string) int32 {
	var count int32
	length := len(words)
	for i := 0; i < length; i++ {
		for j := (i + 1); j < length; j++ {
			if words[i] == words[j] {
				count++
			}
		}
	}
	return count
}

//https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem
func FindAnagrams(s string) int32 {
	var anag int32
	length := len(s)
	words := make([]string, length)
	charCount := 1
	for i := 0; i < length-1; i++ {
		arrLength := length - (charCount - 1)
		wordSlice := words[0:arrLength]
		sliceIdx := 0
		for j := sliceIdx; j < arrLength; j++ {
			chars := string(s[j : charCount+j])
			if len(chars) > 1 {
				chars = sortStringByCharacter(chars)
			}
			wordSlice[sliceIdx] = chars
			sliceIdx++
		}
		anag = anag + countAnagrams(wordSlice)
		charCount++
	}
	return anag
}
