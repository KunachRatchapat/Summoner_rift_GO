package controller

import(
	_cardManagingService "github.com/tehdev/summoner-rift-api/pkg/cardShopMange/service"
)

type cardManagingControllermpl struct{
	cardManagingService _cardManagingService.CardManagingService
}

func NewCardManagingControllermpl (
	cardManagingService  _cardManagingService.CardManagingService,
	
) CardManagingController { //interface
	return &cardManagingControllermpl{cardManagingService}
}