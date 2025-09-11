package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/tehdev/summoner-rift-api/entities"
	"gorm.io/gorm"
	
)

type cardShopRepositoryImpl struct{
	db *gorm.DB
	logger echo.Logger
} 

func NewCardShpRepositoryImpl(db *gorm.DB, Logger echo.Logger) CardshopRepository {
	return &cardShopRepositoryImpl{db, Logger} //return this for implement
}

func (r *cardShopRepositoryImpl) Listing() ([]*entities.Card, error){
	cardList := make([]*entities.Card, 0)

	if err := r.db.Find(&cardList).Error; err != nil {
		r.logger.Error("Failed list Cards: %s",err.Error())
		return nil, err
	}

	return cardList, nil
}