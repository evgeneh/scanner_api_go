package controllers

import (
	"awesomeProject/models"
	"awesomeProject/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllElms(context *gin.Context) {
	var elms []models.Elm
	models.DB.Find(&elms)

	context.JSON(http.StatusOK, gin.H{"data": elms})
}

func StartScanElms(context *gin.Context) {

	var path = context.Query("path")

	if len(path) > 0 {
		services.ElmScan(path)
		context.JSON(http.StatusOK, gin.H{"data": "start"})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": "PATH parameter not found"})
	}

}
