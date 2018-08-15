package maps

func CountTriplets(arr []int64, r int64) int64 {

	freq := make(map[int64]int)  //freq is used to hold count of needed values after this one to complete 2nd part of triplet
	freq2 := make(map[int64]int) //freq2 is used to hold count of needed values to complete triplet

	length := len(arr)
	count := 0

	for i := 0; i < length; i++ {

		_, ok2 := freq2[arr[i]]
		if ok2 { //This value completes freq2[arr[i]] triplets
			count += freq2[arr[i]]
		}

		_, ok := freq[arr[i]]
		if ok { //This value is valid as 2° part of fre[arr[i]] triplets
			freq2[arr[i]*r] += freq[arr[i]]
		}

		freq[arr[i]*r]++ //"Push-up" this value as possible triplet start
	}

	return int64(count)
}
