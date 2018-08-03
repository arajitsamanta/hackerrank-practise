package arrays

//Gcd of two numbers using recursion
func Gcd(a, b int32) int32 {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

//https://www.hackerrank.com/challenges/ctci-array-left-rotation/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
func rotateLeft(arr []int32, d int32) []int32 {

	var i, n int32
	n = int32(len(arr))

	for i = 0; i < Gcd(n, d); i++ {
		temp := arr[i]
		j := i
		for {
			k := j + d
			if k >= n {
				k = k - n
			}
			if k == i {
				break
			}
			arr[j] = arr[k]
			j = k
		}
		arr[j] = temp
	}

	return arr
}
