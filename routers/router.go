package routers

import (
	"github.com/esmira23/go-postgresql-docker/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/post_data", controllers.PostData)
		api.GET("/all", controllers.GetAll)
		api.GET("/transaction/:id", controllers.GetByTransactionId)
		api.GET("/terminal", controllers.GetByTerminalId)
		api.GET("/status/:status", controllers.GetByStatus)
		api.GET("/payment_type/:type", controllers.GetByPaymentType)
		api.GET("/date", controllers.GetByDatePost)
		api.GET("/payment_narrative/:narrative", controllers.GetByPaymentNarrative)

	}
	return router
}
