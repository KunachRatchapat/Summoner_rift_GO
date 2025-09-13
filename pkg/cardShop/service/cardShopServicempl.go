package service

import (

	_cardShopRepository "github.com/tehdev/summoner-rift-api/pkg/cardShop/repository"
	_cardShopModel 		"github.com/tehdev/summoner-rift-api/pkg/cardShop/model"
						
)

type cardShopServiceImpl struct {
	cardShopRepository _cardShopRepository.CardshopRepository //variable file/nameInterface
}

func NewCardShopServiceImpl(
	cardShopRepository _cardShopRepository.CardshopRepository,
	)CardShopService {
		return &cardShopServiceImpl{cardShopRepository}
}

func (s *cardShopServiceImpl) Listing() ([] *_cardShopModel.Card, error) {
	cardList, err := s.cardShopRepository.Listing()
	if err != nil{
		return nil, err
	}

	cardModelList := make([]*_cardShopModel.Card,0)
	for _, card := range cardList{
		cardModelList = append(cardModelList, card.ToCardModel())
	}

	return  cardModelList, nil
}
