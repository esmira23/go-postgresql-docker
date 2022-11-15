package csvparser

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/esmira23/go-postgresql-docker/models"
)

func toInt(value string) int {
	val, _ := strconv.Atoi(value)
	return val
}

func toFloat(value string) float64 {
	val, _ := strconv.ParseFloat(value, 64)
	return val
}

func csvParser() []models.Data {

	csvFile, err := os.Open("csvparser/example.csv")

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(csvFile)
	allData, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	var data []models.Data
	for _, row := range allData[1:] {
		data = append(data, models.Data{
			TransactionId:      toInt(row[0]),
			RequestId:          toInt(row[1]),
			TerminalId:         toInt(row[2]),
			PartnerObjectId:    toInt(row[3]),
			AmountTotal:        toFloat(row[4]),
			AmountOriginal:     toFloat(row[5]),
			CommissionPS:       toFloat(row[6]),
			CommissionClient:   toFloat(row[7]),
			CommissionProvider: toFloat(row[8]),
			DateInput:          row[9],
			DatePost:           row[10],
			Status:             row[11],
			PaymentType:        row[12],
			PaymentNumber:      row[13],
			ServiceId:          toInt(row[14]),
			Service:            row[15],
			PayeeId:            toInt(row[16]),
			PayeeName:          row[17],
			PayeeBankMfo:       toInt(row[18]),
			PayeeBankAccount:   row[19],
			PaymentNarrative:   row[20],
		})
	}

	return data
}

func GetCSVData() []models.Data {
	return csvParser()
}
