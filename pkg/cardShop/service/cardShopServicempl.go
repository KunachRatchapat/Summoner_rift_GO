package service

import (

	_cardShopRepository "github.com/tehdev/summoner-rift-api/pkg/cardShop/repository"
)

type cardShopServiceImpl struct {
	cardShopRepository _cardShopRepository.CardshopRepository //variable file/nameInterface
}

func NewCardShpRepositoryImpl(
	cardShopRepository _cardShopRepository.CardshopRepository,)cardShopService {
		return &cardShopServiceImpl{cardShopRepository}
}

ื