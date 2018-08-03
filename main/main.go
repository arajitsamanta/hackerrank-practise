package main

import (
	"hackerrank-practise/arrays"
)

func main() {
	//result := arrays.Gcd(13, 3)
	//fmt.Println("result:", result)
	a := []int32{1, 2, 5, 3, 7, 8, 6, 4}

	//a := []int32{2, 1, 5, 3, 4}

	b := []int32{2, 5, 1, 3, 4}

	arrays.MinimumBribe(a)
	arrays.MinimumBribe(b)
}
