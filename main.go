package main

import (
	"awesomeProject/controllers"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()

	models.ConnectDB()

	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Is work!"})
	})

	route.GET("/elms", controllers.GetAllElms)

	route.GET("/scan", controllers.StartScanElms)

	route.Run()
}
