package maps

import (
	"fmt"
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

//https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem
func FindAnagrams(s string) int32 {
	var  anaEachItr, anag int32
	words := make(map[string]int)

	length := len(s)

	//start := 0
	charCount := 1
	for charCount < length {
		for i := 0; i < length; i++ {
			//right := charCount + j
			anaEachItr = 0
			for j := i; j < length && (charCount+j) <= length; j++ {
				//fmt.Printf("%d %d", j, (charCount + j))
				chars := string(s[j : charCount+j])

				fmt.Print(string(" " + chars + " "))

				val, ok := words[chars]
				//fmt.Printf(" ok %v val %d ", ok, val)
				if ok && val != 0 {
					/*if anaEachItr == 0 {
						words[chars] = int(anaEachItr)
					} else {
						words[chars]++
						anaEachItr++
					}*/
					/*if anaEachItr == 0 {
						words[chars] = int(anaEachItr + 1)
					} else {

					}*/

					if i==j {
						//reset for second iteration
						words[chars] = 1
					}else{
						
						words[chars]++
						anaEachItr++
					}

					
					/*else {
						anaEachItr++
						words[chars] = words[chars] + val
					}*/
					/*if val > 1 {
						anaEachItr++
					}*/
					//fmt.Printf("\n%v", words)
				} else {
					sortedChars := sortStringByCharacter(chars)
					_, okSorted := words[sortedChars]
						if okSorted {
							anaEachItr++
					}
					words[chars]++
				}

				/*if j == (charCount+j)+charCount {
					words[chars] = 0
				}*/
			}

			anag = anag + anaEachItr
			fmt.Printf("\n anag:%d anagEach:%d", anag, anaEachItr)
			fmt.Println()

			//resert map
			for k,v:=range words{
				v=0
				words[k]=v
			}
		}

		//fmt.Println()

		charCount++
	}

	//fmt.Printf("\n%v", words)

	/*
		for key, val := range anagrams {
			fmt.Printf("\nkey,val %s,%d", key, val)
			lengthOfEachString := len(key)

			if lengthOfEachString > 1 {
				sortedChars := sortStringByCharacter(key)
				valSorted, ok := anagrams[sortedChars]
				if ok {
					if valSorted == 1 {
						count++
					} else {
						count = count + int32(val)
					}
				}
				//fmt.Printf("\n%s %d", chars, val)
			} else {
				count = count + int32(val-1)
			}

			fmt.Printf("\ncount:%d", count)

		}
	*/
	fmt.Printf("\nFinal count:%d", anag)
	return anag
}
