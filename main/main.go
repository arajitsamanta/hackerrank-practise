package main

import (
	"bufio"
	"fmt"
	"hackerrank-practise/maps"
	"io"
	"log"
	"os"
	"strconv"
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
	/*reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//file, err := os.Open("c:/Temp/input12.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewReaderSize(file, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(1000000)

	var queries [][]int32
	for i := 0; i < int(q); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

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
	}

	ans := maps.FreqQuery(queries)

	for i, ansItem := range ans {
		fmt.Fprintf(writer, "%d", ansItem)

		if i != len(ans)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()*/

	total := 0 // count lines
	//begin := time.Now()

	f, err := os.OpenFile("c:/Temp/input12.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var queries [][]int32
	for sc.Scan() {
		row := sc.Text()
		//fmt.Println(row)

		queriesRowTemp := strings.Split(strings.TrimRight(row, " \t\r\n"), " ")

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
	}

	begin := time.Now()
	ans := maps.FreqQuery(queries)

	/*for i, ansItem := range ans {
		fmt.Printf("i %d, item %d", i, ansItem)
		fmt.Println()
	}*/

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
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

	log.Printf("\ntime_used: %v", time.Since(begin).Seconds())

	//scanFile()
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

func scanFile() error {
	total := 0 // count lines
	begin := time.Now()

	f, err := os.OpenFile("c:/Temp/input12.txt", os.O_RDONLY, os.ModePerm)
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
