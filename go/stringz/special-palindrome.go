package stringz

type occurNode struct {
	char  string
	count int64
}

func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func isSpecialPalindrome(str string, n int32) bool {
	var left, right int32
	left = 0
	right = n - 1
	temp := str[0]
	for right > left {
		if temp != str[left] || temp != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

//https://www.hackerrank.com/challenges/special-palindrome-again/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=strings
func substrCountBruteForce(n int32, s string) int64 {
	var i, subStrSize int32
	var count int64
	count = 0
	for subStrSize = 2; subStrSize <= n; subStrSize++ {
		for i = 0; i < n-1; i++ {
			if i+subStrSize <= n {
				temp := s[i : i+subStrSize]
				if isSpecialPalindrome(temp, subStrSize) {
					count++
				}
			}
		}
	}
	return (count + int64(n))
}

func SubstrCount(n int32, s string) int64 {
	var result, count int64
	var i int32

	//Initiliaze an array of occurence struct/node
	occurList := []occurNode{}
	temp := s[0]
	count = 1

	//Put count of same consecutive character in a list. For s="aabbc" , list will be [{a:2} {b:2} {c:1}]
	for i = 1; i < n; i++ {
		if s[i] != temp {
			on := occurNode{char: string(temp), count: count}
			occurList = append(occurList, on)
			count = 0
			temp = s[i]
		}
		count++
	}
	//once the loop ends append count of last char to the list
	on := occurNode{char: string(temp), count: count}
	occurList = append(occurList, on)

	//case 1:  All of the characters are the same, e.g. aaa.
	//To find out no of possible substring we can iterate over the list and add k*(k+1)/2 (total number of substring possible)
	//to the result. Here K  is count of Continuous same char).
	for _, node := range occurList {
		result += node.count * (node.count + 1) / 2
	}

	//case 2: All characters except the middle one are the same, e.g. aadaa.
	//we find all chars that occur once, have neighbouring characters and these neighbouring characters are the same character.
	//For each of these cases we add to our result the minimum of both neighbours. Those match case 2 and that’s it.
	listLen := int32(len(occurList) - 1)
	for i = 1; i < listLen; i++ {
		if occurList[i].count == 1 && occurList[i-1].char == occurList[i+1].char {
			result += min(occurList[i-1].count, occurList[i+1].count)
		}
	}
	return result
}
