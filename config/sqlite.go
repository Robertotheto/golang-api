package config

import (
	"github.com/Robertotheto/golang-api/schemas"
	"gorm.io/gorm"
	"os"
)
import "gorm.io/driver/sqlite"

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	// Check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		// Create the database file and directory
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
	db, err := gorm.Open(sqlite.Open("./db/main.db"), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %s", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %s", err)
		return nil, err
	}
	return db, nil
}
