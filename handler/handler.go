package handler

import (
	"github.com/guilhermemcardoso/go-opportunities-api/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQL()
}
