package greedyalgo

import (
	"sort"
)

//https://www.hackerrank.com/challenges/luck-balance/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=greedy-algorithms
func luckBalance(k int32, contests [][]int32) int32 {

	var (
		temp       []int32
		luck       int32
		i          int32
		mustWin    int32
		luckToFlip int32
		n          int32
	)

	n = int32(len(contests))

	//There is no point in winning unimportant contests, so we can just lose all of them and get their luck.So add all luck first
	for i = 0; i < n; i++ {
		if contests[i][1] == 1 {
			temp = append(temp, contests[i][0])
		}
		luck += contests[i][0]
	}

	//sort temp array, complexity is O(n*log(n))
	sort.SliceStable(temp, func(i, j int) bool {
		return temp[i] < temp[j]
	})

	//Let's denote the number of important contests as . If , we can just lose all of them and add their luck to our luck balance. If , we want to win the  preliminary competitions
	//having the smallest luck value. To do this, we sort all the luck values and subtract the first  values from our luck balance.
	mustWin = int32(len(temp)) - k
	luckToFlip = 0
	for i = 0; i < mustWin; i++ {
		luckToFlip += temp[i]
	}

	return (luck - 2*luckToFlip)
}
