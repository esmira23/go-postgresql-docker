package main

import (
	"github.com/esmira23/go-postgresql-docker/config"
	"github.com/esmira23/go-postgresql-docker/csvparser"

	_ "github.com/esmira23/go-postgresql-docker/docs"
	"github.com/esmira23/go-postgresql-docker/routers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Go API
// @version 1.0
// @description API Server

// @host localhost:8000
// @BasePath /api

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	gin.SetMode(gin.ReleaseMode)
	router := routers.SetUpRouter()
	db := config.ConnectDB()

	config.DB = db

	tx := db.MustBegin()
	csvparser.InserData(tx)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8000")

}
