package controller

import _cardShopService "github.com/tehdev/summoner-rift-api/pkg/cardShop/service"

type cardShopControllerImpl struct{ //strut private
	cardShopService _cardShopService.CardShopService
} 

// constructor function
func NewCardShopControllerImpl(
	cardShopService _cardShopService.CardShopService,
) CardShopController {
	return &cardShopControllerImpl{cardShopService}
}


