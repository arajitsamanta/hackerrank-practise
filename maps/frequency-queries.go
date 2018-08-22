package maps

// https://www.hackerrank.com/challenges/frequency-queries/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps
func FreqQuery(queries [][]int32) []int32 {
	dataMap := make(map[int32]int32)
	freqMap := make(map[int32]int32)
	length := len(queries)
	resultArr := make([]int32, length)
	resultArrIdx := 0

	for i := 0; i < length; i++ {
		row := queries[i]
		switch row[0] {
		//insert query
		case 1:
			dataMap[row[1]]++
			freqMap[dataMap[row[1]]]++

			freqMap[dataMap[row[1]]-1]--

			if freqMap[dataMap[row[1]-1]] < 0 {
				freqMap[dataMap[row[1]-1]] = 0
			}
		//delete query
		case 2:
			dataMap[row[1]]--
			freqMap[dataMap[row[1]]]++

			freqMap[dataMap[row[1]]+1]--

			if freqMap[dataMap[row[1]+1]] < 0 {
				freqMap[dataMap[row[1]+1]] = 0
			}

			if dataMap[row[1]] < 0 {
				dataMap[row[1]] = 0
			}
		//search
		case 3:
			val, ok := freqMap[row[1]]
			if ok && val > 0 {
				resultArr[resultArrIdx] = 1
			}
			resultArrIdx++
		}
	}

	return resultArr[:resultArrIdx]
}

func FreqQuerySlow(queries [][]int32) []int32 {
	freqMap := make(map[int32]int32)
	length := len(queries)
	resultArr := make([]int32, length)
	resultArrIdx := 0

	for i := 0; i < length; i++ {
		row := queries[i]
		switch row[0] {
		case 1:
			freqMap[row[1]]++
		case 2:
			v, ok := freqMap[row[1]]
			if ok && v != 0 {
				freqMap[row[1]]--
			}
		case 3:
			for _, v := range freqMap {
				if row[1] == v {
					resultArr[resultArrIdx] = 1
					break
				}
			}
			resultArrIdx++
		}
	}
	return resultArr[:resultArrIdx]
}
