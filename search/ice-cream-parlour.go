package search

import "fmt"

//https://www.hackerrank.com/challenges/ctci-ice-cream-parlor/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=search
func whatFlavors(cost []int32, money int32) {
	var i, n, target int32
	costsHash := make(map[int32]int32)

	n = int32(len(cost))
	//Two sum array problem , complexity is O(n)
	for i = 0; i < n; i++ {
		target = money - cost[i]
		//if target is found in the array, then print the ouput and break the loop
		_, ok := costsHash[target]
		if ok {
			fmt.Println(costsHash[target], (i + 1))
			break
		}
		costsHash[cost[i]] = i + 1
	}
}
