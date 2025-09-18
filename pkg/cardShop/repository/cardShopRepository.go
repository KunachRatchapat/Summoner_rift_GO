package repository

import (
	"github.com/tehdev/summoner-rift-api/entities"
	_cardShopModel "github.com/tehdev/summoner-rift-api/pkg/cardShop/model"
)


type CardshopRepository interface {
	Listing(cardFilter *_cardShopModel.CardFilter) ([]*entities.Card, error)
	Counting(cardFilter *_cardShopModel.CardFilter) (int64, error)
	FindByID(cardID uint64) (*entities.Card, error)
}