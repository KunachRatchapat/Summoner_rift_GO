package service

import(
	_cardManagingModel "github.com/tehdev/summoner-rift-api/pkg/cardShopMange/model"
	_cardShopModel 	"github.com/tehdev/summoner-rift-api/pkg/cardShop/model"
)

type CardManagingService interface{
	Creating(cardCreatingReq *_cardManagingModel.CardCreatingReq) (*_cardShopModel.Card, error)
	Editing(cardID  uint64, cardEditingReq *_cardManagingModel.CardEditingReq)(*_cardShopModel.Card, error)
	Archiving(cardID uint64) error

}