package main

import (
	"bufio"
	"fmt"
	"hackerrank-practise/go/stringz"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

//"hackerrank-practise/arrays"

//func main() {
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

//maps.FindAnagrams("ifailuhkqq")

//maps.FindAnagramFast("abba")

//maps.FindAnagrams("abcd")

//maps.FindAnagrams("ifailuhkqq")

//maps.FindAnagrams("kkkk")

//s := "abad"
//fmt.Println(SortStringByCharacter(s))

//input := []int64{1, 3, 9, 9, 27, 81}
//input := []int64{1, 2, 2, 4}
//input := []int64{1, 5, 5, 25, 125}

//input := []int64{55, 55, 55, 55}

//input := []int64{1, 2, 1, 2, 4}

//maps.CountTriplets(input, 2)
/*arr := [][]int32{
		{1, 3},
		{2, 3},
		{3, 2},
		{1, 4},
		{1, 5},
		{1, 5},
		{1, 4},
		{3, 2},
		{2, 4},
		{3, 2},
	}

	maps.FreqQuery(arr)
}*/

func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	/*file, err := os.Open("c:/Temp/input2.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewReaderSize(file, 16*1024*1024)*/

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer stdout.Close()

	//writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	/*qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)*/
	/*
		expenditureTemp := strings.Split(readLine(reader), " ")
		var expenditure []int32

		n := int32(200000)

		for i := 0; i < int(n); i++ {
			expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
			checkError(err)
			expenditureItem := int32(expenditureItemTemp)
			expenditure = append(expenditure, expenditureItem)
		}

		begin := time.Now()

		input := []int32{2, 3, 4, 2, 3, 6, 8, 4, 5}
		result := sorting.ActivityNotifications(input, 5)

		log.Printf("\nFast: time_used: %v", time.Since(begin).Seconds())
		fmt.Printf("\nresult: %d", result)*/

	/*
		ans := maps.FreqQuery(queries)

		for i, ansItem := range ans {
			fmt.Fprintf(writer, "%d", ansItem)

			if i != len(ans)-1 {
				fmt.Fprintf(writer, "\n")
			}
		}

		fmt.Fprintf(writer, "\n")

		writer.Flush()*/

	/*
		f, err := os.OpenFile("c:/Temp/input2.txt", os.O_RDONLY, os.ModePerm)
		if err != nil {
			log.Fatalf("open file error: %v", err)
		}
		defer f.Close()

		sc := bufio.NewScanner(f)
		var queries [][]int32
		var expenditure []int32
		for sc.Scan() {
			row := sc.Text()
			expenditureItemTemp, err := strconv.ParseInt(row, 10, 64)
			checkError(err)
			expenditureItem := int32(expenditureItemTemp)
			expenditure = append(expenditure, expenditureItem)

			/*queriesRowTemp := strings.Split(strings.TrimRight(row, " \t\r\n"), " ")

			var queriesRow []int32
			for _, queriesRowItem := range queriesRowTemp {
				queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
				checkError(err)
				queriesItem := int32(queriesItemTemp)
				queriesRow = append(queriesRow, queriesItem)
			}

			if len(queriesRow) != 2 {
				panic("Bad input")
			}

			queries = append(queries, queriesRow)

			total++
		}

		if err := sc.Err(); err != nil {
			log.Fatalf("scan file error: %v", err)
		}*/

	//ans := maps.FreqQuery(queries)

	/*for i, ansItem := range ans {
		fmt.Printf("i %d, item %d", i, ansItem)
		fmt.Println()
	}*/

	/*stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	for i, ansItem := range ans {
		fmt.Fprintf(writer, "%d", ansItem)

		if i != len(ans)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()

	log.Printf("\ntime_used: %v", time.Since(begin).Seconds())*/

	//scanFile()
	//a := []int32{2, 3, 4, 2, 3, 6, 8, 4, 5}
	//sorting.CountSwaps(a)
	//

	/*input := []int32{7, 5, 3, 1}
	begin := time.Now()
	result := sorting.CountInversions(input)
	fmt.Printf("\ntime_used: %v", time.Since(begin).Seconds())
	fmt.Printf("\nresult: %d", result)*/
	//begin := time.Now()
	result := stringz.SubstrCount(8, "mnonopoo")
	//fmt.Printf("\ntime_used: %v", time.Since(begin).Seconds())
	fmt.Printf("\ncount: %d", result)
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
}

func scanFileLineByLine() error {
	total := 0 // count lines
	begin := time.Now()

	f, err := os.OpenFile("c:/Temp/input2.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		row := sc.Text()
		fmt.Println(row)
		total++
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return err
	}

	log.Printf("scan file, time_used: %v, lines=%v\n", time.Since(begin).Seconds(), total)

	return nil
}
