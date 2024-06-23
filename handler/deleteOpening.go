package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilhermemcardoso/go-opportunities-api/schemas"
)

func DeleteOpeningHandler(ctx *gin.Context) {

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	// find opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}
	// delete opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "failed to delete opening")
		return
	}
	sendSuccess(ctx, "delete-opening", opening)
}
