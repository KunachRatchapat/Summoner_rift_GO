package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/tehdev/summoner-rift-api/entities"
	"gorm.io/gorm"
	_cardShopException "github.com/tehdev/summoner-rift-api/pkg/cardShop/exception"
	_cardShopModel		"github.com/tehdev/summoner-rift-api/pkg/cardShop/model"
	
)

type cardShopRepositoryImpl struct{
	db *gorm.DB
	logger echo.Logger
} 


func NewCardShpRepositoryImpl(db *gorm.DB, Logger echo.Logger) CardshopRepository {
	return &cardShopRepositoryImpl{db, Logger} //return this for implement
}


func (r *cardShopRepositoryImpl) Listing(cardFilter *_cardShopModel.CardFilter) ([]*entities.Card, error){
	cardList := make([]*entities.Card, 0)

	//Select * from
	query := r.db.Model(&entities.Card{}).Where("is_archive = ?", false) 

	if cardFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+cardFilter.Name+"%")
	}

	if cardFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+cardFilter.Description+"%")
	}

	offset := int((cardFilter.Page-1) * cardFilter.Size)
	limit := int(cardFilter.Size)

	if err := query.Offset(offset).Limit(limit).Find(&cardList).Error; err != nil {
		r.logger.Error("Failed list Cards: %s",err.Error())
		return nil, &_cardShopException.CardListing{}
	}

	return cardList, nil
}

func (r *cardShopRepositoryImpl) Counting(cardFilter *_cardShopModel.CardFilter) (int64, error){

	//Select * from
	query := r.db.Model(&entities.Card{}).Where("is_archive = ?", false) 

	if cardFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+cardFilter.Name+"%")
	}

	if cardFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+cardFilter.Description+"%")
	}

	var count int64

	
	if err := query.Count(&count).Error; err != nil {
		r.logger.Error("Counting Card Failed: %s",err.Error())
		return -1, &_cardShopException.CardCounting{}
	}

	return count, nil
}