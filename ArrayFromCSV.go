package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	outputArray := ArrayFromCSV("./sample.csv")
	fmt.Println(outputArray)
}

func ArrayFromCSV(filePath string) []float64 {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	var data float64
	var datas []float64

	for _, record := range records[0] {

		data, _ = strconv.ParseFloat(record, 64)
		datas = append(datas, data)
	}

	return datas
}
