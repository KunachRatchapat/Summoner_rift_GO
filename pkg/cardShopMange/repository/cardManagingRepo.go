package repository

import (
	"github.com/tehdev/summoner-rift-api/entities"
	_cardManagingModel"github.com/tehdev/summoner-rift-api/pkg/cardShopMange/model"
)


type CardManagingRepository interface {
	Creating(cardEntity		*entities.Card) (*entities.Card, error)
	Editing(cardID uint64, cardEditing *_cardManagingModel.CardEditingReq) (uint64, error)
	Archiving(cardID uint64) error
}