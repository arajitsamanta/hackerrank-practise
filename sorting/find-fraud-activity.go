package sorting

// https://www.hackerrank.com/challenges/fraudulent-activity-notifications/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=sorting
//Below approach use counting sort and prefix sum array approach to find median of an array.
//O(n*k) - n: Size of input array, k: max of array
func ActivityNotifications(expenditure []int32, d int32) int32 {
	var i, n, median, count int32
	n = int32(len(expenditure))
	freq := make([]int32, 200)
	firstTime := true
	var popElement int32
	for index := d; index < n; index++ {

		if firstTime {
			firstTime = false
			for i = index - d; i <= index-1; i++ {
				freq[expenditure[i]]++
			}
		} else {
			freq[popElement]--
			freq[expenditure[index-1]]++
		}

		median = getMedian(freq, d)
		if d%2 == 0 {
			if expenditure[index] >= median {
				count++
			}
		} else {
			if expenditure[index] >= 2*median {
				count++
			}
		}
		popElement = expenditure[index-d]
	}
	return count
}

func getMedian(freq []int32, d int32) int32 {
	prefixSum := make([]int32, 200)
	prefixSum[0] = freq[0]

	for i := 1; i < 200; i++ {
		prefixSum[i] = prefixSum[i-1] + freq[i]
	}

	var median, a, b, first, second, i int32

	if d%2 == 0 {
		first = d / 2
		second = first + 1
		i = 0
		for ; i < 200; i++ {
			if first <= prefixSum[i] {
				a = i
				break
			}
		}
		for ; i < 200; i++ {
			if second <= prefixSum[i] {
				b = i
				break
			}
		}

	} else {
		first = d/2 + 1
		for i = 0; i < 200; i++ {
			if first <= prefixSum[i] {
				a = i
				break
			}
		}
	}

	median = a + b

	return median
}
