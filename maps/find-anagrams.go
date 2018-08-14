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

func hash(freq []int) int64 {

	var ret, t int64
	t = 1
	for i := 0; i < 26; i++ {
		ret += t * int64(freq[i])
		t = t * 1000003
	}
	return ret
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

func FindAnagramFast(s string) int32 {

	var count int
	freqMap := make(map[int64]int)
	length := len(s)

	for i := 0; i < length; i++ {

		freq := make([]int, 26)
		for j := i; j < length; j++ {
			freq[s[j]-'a']++
			freqMap[hash(freq)]++
		}
	}

	for _, v := range freqMap {
		count = count + (v * (v - 1) / 2)
	}

	return int32(count)
}
