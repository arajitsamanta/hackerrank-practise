package search

//https://www.hackerrank.com/challenges/pairs/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=search
//Runtime c
func pairs(k int32, arr []int32) int32 {
	var target, count int32
	arrHash := make(map[int32]int32)
	n := len(arr)

	//build the hash table first - O(n)
	for i := 0; i < n; i++ {
		arrHash[arr[i]]++
	}

	//Run through the array and find target - O(n)
	for i := 0; i < n; i++ {
		target = arr[i] + k
		_, ok := arrHash[target]
		if ok {
			count++
		}
	}

	return count
}
