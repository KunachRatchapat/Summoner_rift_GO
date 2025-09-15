package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type cardManagingRepositoryMpl struct{
	db *gorm.DB
	logger echo.Logger
}

func NewCardManagingRepositorympl (db *gorm.DB, Logger echo.Logger) *cardManagingRepositoryMpl {
	return &cardManagingRepositoryMpl{
		db: db,
		logger: Logger,
	}
}