package maps

import (
	"strings"
)

//https://www.hackerrank.com/challenges/two-strings/problem
func TwoString(s1 string, s2 string) string {

	letters := "abcdefghijklmnopqrstuvwxyz"

	result := "NO"
	for _, v := range letters {
		if strings.Index(s1, string(v)) > -1 && strings.Index(s2, string(v)) > -1 {
			result = "YES"
		}
	}
	return result
}
