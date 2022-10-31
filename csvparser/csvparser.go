package csvparser

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/esmira23/go-postgresql-docker/models"
	"github.com/jmoiron/sqlx"
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
	rows, _ := reader.ReadAll()

	var data []models.Data
	for _, row := range rows[1:] {
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

func InserData(tx *sqlx.Tx) {

	csvdata := csvParser()

	for i := 0; i < len(csvdata); i++ {
		tx.MustExec("INSERT INTO example VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21) ON CONFLICT (transaction_id) DO NOTHING", csvdata[i].TransactionId, csvdata[i].RequestId, csvdata[i].TerminalId, csvdata[i].PartnerObjectId, csvdata[i].AmountTotal, csvdata[i].AmountOriginal, csvdata[i].CommissionPS, csvdata[i].CommissionClient, csvdata[i].CommissionProvider, csvdata[i].DateInput, csvdata[i].DatePost, csvdata[i].Status, csvdata[i].PaymentType, csvdata[i].PaymentNumber, csvdata[i].ServiceId, csvdata[i].Service, csvdata[i].PayeeId, csvdata[i].PayeeName, csvdata[i].PayeeBankMfo, csvdata[i].PayeeBankAccount, csvdata[i].PaymentNarrative)
	}

	tx.Commit()
}

func GetCSVData() []models.Data {
	return csvParser()
}
