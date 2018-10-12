package search

// https://www.hackerrank.com/challenges/minimum-time-required/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=search
// RUntime Complexity: O(k*n) +  O(n) = O(k*n)
func MinTimeBruteForce(machines []int64, goal int64) int64 {

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

/*
This problem can be solved using binary search over the answer. Since we have the low=1 and high=1e18 and  pointers, we can apply binary search.
Now, at every iteration we check if we can complete our task in  days by traversing the array and summing up the total items produced by
each machine. If yes, then we move towards left else we move towards right (Since, we need to minimize the answer).
*/
func minTime(machines []int64, goal int64) int64 {

	var i, n, low, high, mid, ans, done int64

	//goal is atleast , then we need a minimum of  day to complete the task.
	low = 1
	//maximum of  days (when goal is  and the number of machine is  and it's rate is ).
	high = 1e18

	n = int64(len(machines))
	for low <= high {
		mid = (low + high) / 2
		done = 0
		for i = 0; i < n; i++ {
			done += mid / machines[i]
			if done >= goal {
				break
			}

		}
		if done >= goal {
			high = mid - 1
			ans = mid
		} else {
			low = mid + 1
		}

	}
	return ans
}
