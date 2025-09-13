package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_cardShopService "github.com/tehdev/summoner-rift-api/pkg/cardShop/service"
)

type cardShopControllerImpl struct{ //strut private
	cardShopService _cardShopService.CardShopService
} 

// constructor function
func NewCardShopControllerImpl(
	cardShopService _cardShopService.CardShopService,
) CardShopController {
	return &cardShopControllerImpl{cardShopService}
}

func (c *cardShopControllerImpl) Listing(pctx echo.Context) error {
	cardModelList, err := c.cardShopService.Listing()
	if err != nil{
		return  pctx.String(http.StatusInternalServerError,err.Error())
	}

	return pctx.JSON(http.StatusOK, cardModelList)
}


