package controllers

import (
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"

	"github.com/esmira23/go-postgresql-docker/config"
	"github.com/esmira23/go-postgresql-docker/csvparser"
	"github.com/esmira23/go-postgresql-docker/models"
	"github.com/gin-gonic/gin"
)

// @Summary Post All
// @Tags Upload Data
// @Description Post all data
// @ID post
// @Produce json
// @Success 200 {object} models.ErrorMsg
// @Failure 500 {object} models.ErrorMsg
// @Router /post_data [post]
func PostData(c *gin.Context) {

	config.ConnectDB()
	
	defer func(DB *sqlx.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(config.DB)

	tx := config.DB.MustBegin()

	csvdata := csvparser.CSVparser()

	for i := 1; i < len(csvdata); i++ {
		tx.MustExec("INSERT INTO example VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21) ON CONFLICT (transaction_id) DO NOTHING", csvparser.ToInt(csvdata[i][0]), csvparser.ToInt(csvdata[i][1]), csvparser.ToInt(csvdata[i][2]), csvparser.ToInt(csvdata[i][3]), csvparser.ToFloat(csvdata[i][4]), csvparser.ToFloat(csvdata[i][5]), csvparser.ToFloat(csvdata[i][6]), csvparser.ToFloat(csvdata[i][7]), csvparser.ToFloat(csvdata[i][8]), csvdata[i][9], csvdata[i][10], csvdata[i][11], csvdata[i][12], csvdata[i][13], csvparser.ToInt(csvdata[i][14]), csvdata[i][15], csvparser.ToInt(csvdata[i][16]), csvdata[i][17], csvparser.ToInt(csvdata[i][18]), csvdata[i][19], csvdata[i][20])
	}

	err := tx.Commit()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, models.ErrorMsg{
			Message: "data uploaded",
		})
	}

}

// @Summary Get All
// @Tags Get Data
// @Description Get all data
// @ID all
// @Produce json
// @Success 200 {object} models.Data
// @Failure 404 {object} models.ErrorMsg
// @Failure 500 {object} map[string][]string
// @Router /all [get]
func GetAll(c *gin.Context) {

	config.ConnectDB()
	defer config.DB.Close()

	items := []models.Data{}

	err := config.DB.Select(&items, "SELECT * FROM example")

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else if len(items) == 0 {
		c.JSON(http.StatusNotFound, models.ErrorMsg{
			Message: "data not found",
		})
	} else {
		c.JSON(http.StatusOK, items)
	}

}

// @Summary Get By Transaction Id
// @Tags Get Data
// @Description get data by transaction id
// @ID transaction-id
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} models.Data
// @Failure 404 {object} models.ErrorMsg
// @Failure 500 {object} map[string][]string
// @Router /transaction/{id} [get]
func GetByTransactionId(c *gin.Context) {

	config.ConnectDB()
	defer config.DB.Close()

	items := []models.Data{}
	param := c.Param("id")

	err := config.DB.Select(&items, "SELECT * FROM example WHERE transaction_id = $1", param)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else if len(items) == 0 {
		c.JSON(http.StatusNotFound, models.ErrorMsg{
			Message: "data not found",
		})
	} else {
		c.JSON(http.StatusOK, items)
	}

}

// @Summary Get By Terminal Id
// @Tags Get Data
// @Description Get data by terminal id
// @ID terminal-id
// @Produce json
// @Param id query string true "Terminal ID. Pattern: id_1,id_2,..,id_n"
// @Success 200 {object} models.Data
// @Failure 404 {object} models.ErrorMsg
// @Failure 500 {object} map[string][]string
// @Router /terminal [get]
func GetByTerminalId(c *gin.Context) {

	config.ConnectDB()
	defer config.DB.Close()

	items := []models.Data{}
	param := c.Query("id")

	selectquery := "SELECT * FROM example WHERE terminal_id IN (" + param + ")"
	err := config.DB.Select(&items, selectquery)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else if len(items) == 0 {
		c.JSON(http.StatusNotFound, models.ErrorMsg{
			Message: "data not found",
		})
	} else {
		c.JSON(http.StatusOK, items)
	}

}

// @Summary Get By Status
// @Tags Get Data
// @Description get data by status
// @ID status
// @Produce json
// @Param status path string true "Status (accepted/declined)"
// @Success 200 {object} models.Data
// @Failure 404 {object} models.ErrorMsg
// @Failure 500 {object} map[string][]string
// @Router /status/{status} [get]
func GetByStatus(c *gin.Context) {

	config.ConnectDB()
	defer config.DB.Close()

	items := []models.Data{}
	param := c.Param("status")

	err := config.DB.Select(&items, "SELECT * FROM example WHERE status = $1", param)
	defer config.DB.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else if len(items) == 0 {
		c.JSON(http.StatusNotFound, models.ErrorMsg{
			Message: "data not found",
		})
	} else {
		c.JSON(http.StatusOK, items)
	}

}

// @Summary Get By Payment Type
// @Tags Get Data
// @Description get data by payment type
// @ID payment-type
// @Produce json
// @Param type path string true "Payment Type (card/cash)"
// @Success 200 {object} models.Data
// @Failure 404 {object} models.ErrorMsg
// @Failure 500 {object} map[string][]string
// @Router /payment_type/{type} [get]
func GetByPaymentType(c *gin.Context) {

	config.ConnectDB()
	defer config.DB.Close()

	items := []models.Data{}
	param := c.Param("type")

	err := config.DB.Select(&items, "SELECT * FROM example WHERE payment_type = $1", param)
	defer config.DB.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else if len(items) == 0 {
		c.JSON(http.StatusNotFound, models.ErrorMsg{
			Message: "data not found",
		})
	} else {
		c.JSON(http.StatusOK, items)
	}

}

// @Summary Get By Post Date
// @Tags Get Data
// @Description Get data by date of post
// @ID date
// @Produce json
// @Param from query string true "Date from YYYY-MM-DD"
// @Param to query string true "Date to YYYY-MM-DD"
// @Success 200 {object} models.Data
// @Failure 404 {object} models.ErrorMsg
// @Failure 500 {object} map[string][]string
// @Router /date [get]
func GetByDatePost(c *gin.Context) {

	config.ConnectDB()
	defer config.DB.Close()

	items := []models.Data{}
	from := c.Query("from")
	to := c.Query("to")

	err := config.DB.Select(&items, "SELECT * FROM example WHERE date_post::text BETWEEN $1 AND $2", from, to)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else if len(items) == 0 {
		c.JSON(http.StatusNotFound, models.ErrorMsg{
			Message: "data not found",
		})
	} else {
		c.JSON(http.StatusOK, items)
	}

}

// @Summary Get By Payment Narrative
// @Tags Get Data
// @Description Get data by payment narrative
// @ID payment-narrative
// @Produce json
// @Param narrative path string true "Partially specified payment narrative"
// @Success 200 {object} models.Data
// @Failure 404 {object} models.ErrorMsg
// @Failure 500 {object} map[string][]string
// @Router /payment_narrative/{narrative} [get]
func GetByPaymentNarrative(c *gin.Context) {

	config.ConnectDB()
	defer config.DB.Close()

	items := []models.Data{}
	param := c.Param("narrative")

	err := config.DB.Select(&items, "SELECT * FROM example WHERE payment_narrative ILIKE $1", "%"+param+"%")
	defer config.DB.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else if len(items) == 0 {
		c.JSON(http.StatusNotFound, models.ErrorMsg{
			Message: "data not found",
		})
	} else {
		c.JSON(http.StatusOK, items)
	}

}
