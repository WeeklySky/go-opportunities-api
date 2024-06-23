package config

import (
	"os"

	"github.com/guilhermemcardoso/go-opportunities-api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	// Check if DB exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("DB does not exist. Creating...")
		// Create DB file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// Create DB and connect to it
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("failed to connect database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("failed to automigrate schema: %v", err)
		return nil, err
	}

	return db, nil
}
