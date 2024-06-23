package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilhermemcardoso/go-opportunities-api/schemas"
)

func ListOpeningsHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ListOpeningsHandler",
	})

	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "failed to get openings")
		return
	}
	sendSuccess(ctx, "list-openings", openings)
}
