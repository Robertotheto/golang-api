package handler

import (
	"github.com/Robertotheto/golang-api/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	db = config.GetSQLite()
	logger = config.GetLogger("handler")
}
