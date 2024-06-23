package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilhermemcardoso/go-opportunities-api/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {

	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("failed to validate request: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("failed to create opening: %v", err)
		sendError(ctx, http.StatusInternalServerError, "failed to create opening")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
