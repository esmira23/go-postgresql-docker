package csvparser

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func ToInt(value string) int {
	val, _ := strconv.Atoi(value)
	return val
}

func ToFloat(value string) float64 {
	val, _ := strconv.ParseFloat(value, 64)
	return val
}

func CSVparser() [][]string {

	csvFile, err := os.Open("csvparser/example.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	allData, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	return allData

}
