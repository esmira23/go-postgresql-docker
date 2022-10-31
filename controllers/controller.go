package controllers

import (
	"github.com/esmira23/go-postgresql-docker/config"
	"github.com/esmira23/go-postgresql-docker/models"
	"github.com/gin-gonic/gin"
)

// @Summary Get All
// @Tags EndPoints
// @Description Get all data
// @ID all
// @Produce json
// @Success 200 {object} models.Data
// @Failure 404 {object} models.ErrorMsg
// @Router /all [get]
func GetAll(c *gin.Context) {
	items := []models.Data{}

	err := config.DB.Select(&items, "SELECT * FROM example")

	if len(items) == 0 || err != nil {
		c.JSON(404, gin.H{"message": "data not found"})
	} else {
		c.JSON(200, items)
	}
}

// @Summary Get By Transaction Id
// @Tags EndPoints
// @Description get data by transaction id
// @ID transaction-id
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} models.Data
// @Failure 404 {object} models.ErrorMsg
// @Router /transaction/{id} [get]
func GetByTransactionId(c *gin.Context) {
	items := []models.Data{}
	param := c.Param("id")
	err := config.DB.Select(&items, "SELECT * FROM example WHERE transaction_id = $1", param)

	if len(items) == 0 || err != nil {
		c.JSON(404, gin.H{"message": "data not found"})
	} else {
		c.JSON(200, items)
	}
}

// @Summary Get By Terminal Id
// @Tags EndPoints
// @Description Get data by terminal id
// @ID terminal-id
// @Produce json
// @Param id query string true "Terminal ID. Pattern: id_1,id_2,..,id_n"
// @Success 200 {object} models.Data
// @Failure 404 {object} models.ErrorMsg
// @Router /terminal [get]
func GetByTerminalId(c *gin.Context) {
	items := []models.Data{}
	param := c.Query("id")

	selectquery := "SELECT * FROM example WHERE terminal_id IN (" + param + ")"
	err := config.DB.Select(&items, selectquery)

	if len(items) == 0 || err != nil {
		c.JSON(404, gin.H{"message": "data not found"})
	} else {
		c.JSON(200, items)
	}
}

// @Summary Get By Status
// @Tags EndPoints
// @Description get data by status
// @ID status
// @Produce json
// @Param status path string true "Status (accepted/declined)"
// @Success 200 {object} models.Data
// @Failure 404 {object} models.ErrorMsg
// @Router /status/{status} [get]
func GetByStatus(c *gin.Context) {
	items := []models.Data{}
	param := c.Param("status")

	err := config.DB.Select(&items, "SELECT * FROM example WHERE status = $1", param)

	if len(items) == 0 || err != nil {
		c.JSON(404, gin.H{"message": "data not found"})
	} else {
		c.JSON(200, items)
	}
}

// @Summary Get By Payment Type
// @Tags EndPoints
// @Description get data by payment type
// @ID payment-type
// @Produce json
// @Param type path string true "Payment Type (card/cash)"
// @Success 200 {object} models.Data
// @Failure 404 {object} models.ErrorMsg
// @Router /payment_type/{type} [get]
func GetByPaymentType(c *gin.Context) {
	items := []models.Data{}
	param := c.Param("type")

	err := config.DB.Select(&items, "SELECT * FROM example WHERE payment_type = $1", param)

	if len(items) == 0 || err != nil {
		c.JSON(404, gin.H{"message": "data not found"})
	} else {
		c.JSON(200, items)
	}
}

// @Summary Get By Post Date
// @Tags EndPoints
// @Description Get data by date of post
// @ID date
// @Produce json
// @Param from query string true "Date from YYYY-MM-DD"
// @Param to query string true "Date to YYYY-MM-DD"
// @Success 200 {object} models.Data
// @Failure 404 {object} models.ErrorMsg
// @Router /date [get]
func GetByDatePost(c *gin.Context) {
	items := []models.Data{}
	from := c.Query("from")
	to := c.Query("to")

	err := config.DB.Select(&items, "SELECT * FROM example WHERE date_post::text BETWEEN $1 AND $2", from, to)

	if len(items) == 0 || err != nil {
		c.JSON(404, gin.H{"message": "data not found"})
	} else {
		c.JSON(200, items)
	}
}

// @Summary Get By Payment Narrative
// @Tags EndPoints
// @Description Get data by payment narrative
// @ID payment-narrative
// @Produce json
// @Param narrative path string true "Partially specified payment narrative"
// @Success 200 {object} models.Data
// @Failure 404 {object} models.ErrorMsg
// @Router /payment_narrative/{narrative} [get]
func GetByPaymentNarrative(c *gin.Context) {
	items := []models.Data{}
	param := c.Param("narrative")

	err := config.DB.Select(&items, "SELECT * FROM example WHERE payment_narrative ILIKE $1", "%"+param+"%")

	if len(items) == 0 || err != nil {
		c.JSON(404, gin.H{"message": "data not found"})
	} else {
		c.JSON(200, items)
	}
}
