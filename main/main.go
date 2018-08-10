package main

import (
	//"hackerrank-practise/arrays"
	"hackerrank-practise/maps"
)

func main() {
	//result := arrays.Gcd(13, 3)
	//fmt.Println("result:", result)
	//a := []int32{1, 2, 5, 3, 7, 8, 6, 4}

	//a := []int32{2, 1, 5, 3, 4}

	//b := []int32{8, 3, 5, 2, 4, 1, 6}

	//arrays.MinimumBribe(a)
	//arrays.MinimumBribe(b)
	//arrays.MinimumSwaps(b)

	//arrays.MinimumSwaps(b)

	/*a := [][]int32{
		{2, 3, 603},
		{1, 1, 286},
		{4, 4, 882},
	}*/

	/*a := [][]int32{
		{1, 2, 100},
		{2, 5, 100},
		{3, 4, 100},
	}*/

	/*a := [][]int32{
		{2, 6, 8},
		{3, 5, 7},
		{1, 8, 1},
		{5, 9, 15},
	}*/

	//arrays.ArrayManipulation(4, a)

	//mag := []string{"ive", "got", "a", "got", "lovely", "bunch", "of", "coconuts"}
	//notes:= []string{"ive", "got", "some", "coconuts"}

	//mag := []string{"give", "me", "one", "grand", "today", "night"}
	//notes:= []string{"give", "one", "grand", "today"}

	//mag := []string{"two", "times", "three", "is", "not", "four"}
	//notes := []string{"two", "times", "IS", "four"}

	//maps.CheckMagazine(mag, notes)

	//maps.TwoString("hello", "world")

	//maps.FindAnagrams("mom")

	maps.FindAnagrams("abba")

	//maps.FindAnagrams("abcd")

	//maps.FindAnagrams("ifailuhkqq")

	//maps.FindAnagrams("kkkk")

	//s := "abad"
	//fmt.Println(SortStringByCharacter(s))
}

/*
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var queries [][]int32
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != int(3) {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := arrays.ArrayManipulation(n, queries)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}*/
