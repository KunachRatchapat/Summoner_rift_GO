package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/tehdev/summoner-rift-api/databases"
	"github.com/tehdev/summoner-rift-api/entities"
	_cardManagingExcepiton "github.com/tehdev/summoner-rift-api/pkg/cardShopMange/exception"
	_cardManagingModel "github.com/tehdev/summoner-rift-api/pkg/cardShopMange/model"
	
)

type cardManagingRepositoryMpl struct{
	db databases.Database
	logger echo.Logger
}

func NewCardManagingRepositorympl (db databases.Database, Logger echo.Logger) *cardManagingRepositoryMpl {
	return &cardManagingRepositoryMpl{
		db: db,
		logger: Logger,
	}
}

//Function การสร้างการ์ดดด
func (r *cardManagingRepositoryMpl) Creating(cardEntity *entities.Card) (*entities.Card, error){
	card := new(entities.Card) //สร้าง Instance เก็บไว้ใน card	

	if err := r.db.Connect().Create(cardEntity).Scan(card).Error; err != nil{
		r.logger.Errorf("Create card Failed : %s", err.Error())
		return  nil , &_cardManagingExcepiton.CardCreating{}
	}

	return  card, nil
}

//Editing Cards
func (r *cardManagingRepositoryMpl) Editing(cardID uint64, cardEditingReq *_cardManagingModel.CardEditingReq) (uint64, error){
	if err := r.db.Connect().Model(&entities.Card{}).Where(
		"id = ?",cardID,
	).Updates(
		cardEditingReq,
	).Error; err != nil{
		r.logger.Errorf("Editing card failed: %s", err.Error())
		return 0, &_cardManagingExcepiton.CardEditing{}
	} 
	return cardID, nil
}

//Archive
func (r *cardManagingRepositoryMpl) Archiving(cardID uint64) error{
	if err := r.db.Connect().Table(
		"cards",
	).Where(
		"id=?",cardID,
	).Update(
		"is_archive", true,
	).Error; err != nil{
		r.logger.Errorf("Archiving card failed: %s", err.Error())
		return &_cardManagingExcepiton.CardArchving{CardID: cardID}
	}

	return  nil
}
