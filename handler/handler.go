package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "CreateOpeningHandler",
	})
}

func ShowOpeningHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ShowOpeningHandler",
	})
}

func DeleteOpeningHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "DeleteOpeningHandler",
	})
}

func UpdateOpeningHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "UpdateOpeningHandler",
	})
}

func ListOpeningsHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ListOpeningsHandler",
	})
}
