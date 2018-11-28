package stack

//https://www.hackerrank.com/challenges/game-of-two-stacks/problem
func twoStacks(x int32, a []int32, b []int32) int32 {
	var sum, i, j, n, m, count int32
	n = int32(len(a))
	m = int32(len(b))

	for i < n && sum+a[i] <= x {
		sum += a[i]
		i++
	}
	count = i

	for j < m && i >= 0 {
		sum += a[j]
		j++

		for sum > x && i > 0 {
			sum -= a[i]
			i--
		}

		if sum <= x && i+j > count {
			count = i + j
		}

	}
	return count
}
