package service

import (
	_cardManagingRepository "github.com/tehdev/summoner-rift-api/pkg/cardShopMange/repository"
	
)

type cardManagingServicempl struct{
	cardManagingRepository _cardManagingRepository.CardManagingRepository

}

func NewCardManagingServicempl (
	cardManagingRepository _cardManagingRepository.CardManagingRepository,

) CardManagingService{
	return  &cardManagingServicempl{cardManagingRepository}
}

