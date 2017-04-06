package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("NAME\n\tzipf - Produce a cluster of 32-bits integers in a zipf distribution")
		fmt.Println("SYNOPSIS\n\tDictSize DataSize ZipfSkew OutputPath")
		fmt.Println("DESCRIPTION\n\tDictSize: the cardinality of the value domain. And usually DictSize is far less than DataSize")
		fmt.Println("\tDataSize: the number of integers that you want to produce")
		fmt.Println("\tZipfSkew: the tunable parameter of zipf distribution, the bigger it is, the more difference between adjacent ranks")
		fmt.Println("\tFor example:\n\t\t./zipf 1000 10000 1 data")
		return
	}
	var dictSize, dataSize, skew int
	var err error
	if dictSize, err = strconv.Atoi(os.Args[1]); err != nil || dictSize < 1 {
		fmt.Println("DictSize isn't numeric or dictSize is less than 1")
		return
	}
	if dataSize, err = strconv.Atoi(os.Args[2]); err != nil || dataSize < 1 {
		fmt.Println("DataSize isn't numeric or size is less than 1")
		return
	}
	if skew, err = strconv.Atoi(os.Args[3]); err != nil || skew < 1 {
		fmt.Println("Skew isn't numeric or skew is less than 1")
		return
	}
	outputFile := os.Args[4]
	wh, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Errorf("Open file error: %v\n", err)
	}
	defer wh.Close()

	// calculate the frequencies for rank `0` ~ rank `dictSize-1`
	// assign randomly for the numbers
	frequency := make([]float64, dictSize)
	num := make([]int32, dictSize)
	sum := float64(0)

	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < dictSize; i++ {
		frequency[i] = 1.0 / math.Pow(float64(i+1), float64(skew))
		sum += frequency[i]
		num[i] = rand.Int31()
	}
	for i := 0; i < dictSize; i++ {
		frequency[i] /= sum
	}

	rand.Seed(int64(time.Now().Nanosecond()))
	// output the result based on the random rank and the frequency of one rank.
	for i := 0; i < dataSize; i++ {
		var rank int
		var fre float64
		rank = rand.Intn(dictSize)
		fre = frequency[rank]
		for rand.Float64() >= fre {
			rank = rand.Intn(dictSize)
			fre = frequency[rank]
		}
		wh.WriteString(fmt.Sprintf("%d\n", num[rank]))
	}
	fmt.Printf("The #rank1 number is %d, and its frequency is %.5f\n", num[0], frequency[0])
	fmt.Println("Note: If the occurence frequency of the rank1 is high, you can decrease the `ZipfSkew` parameter.")
}
