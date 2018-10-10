package search

// https://www.hackerrank.com/challenges/minimum-time-required/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=search
// RUntime Complexity: O(nlogn) +  O(n) = O(n)
func minTimeBruteForce(machines []int64, goal int64) int64 {

	var produced, dayCount int64

	//Build map O(n)
	machineMap := make(map[int64]int)
	for _, val := range machines {
		machineMap[val]++
	}

	produced = 0
	dayCount = 1

	//O(k*n) , where k is the goal and n is the size of machines
	for produced < goal {
		noOfProducesPerDay := 0
		for k, v := range machineMap {
			if (dayCount % k) == 0 {
				noOfProducesPerDay += v
			}
		}
		produced += int64(noOfProducesPerDay)
		if produced >= goal {
			break
		}
		dayCount++
	}

	return dayCount

}
