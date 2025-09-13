package controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	_cardShopException "github.com/tehdev/summoner-rift-api/pkg/cardShop/exception"
	_cardShopModel "github.com/tehdev/summoner-rift-api/pkg/cardShop/model"
	_cardShopService "github.com/tehdev/summoner-rift-api/pkg/cardShop/service"
	"github.com/tehdev/summoner-rift-api/pkg/custom"
	
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
	cardFileter := new(_cardShopModel.CardFilter)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(cardFileter); err != nil{
		return  custom.Error(pctx,http.StatusInternalServerError,err.Error())
		
	}

	cardModelList, err := c.cardShopService.Listing(cardFileter)
	if err != nil{
			return  custom.Error(pctx, http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusOK, cardModelList)

	
}


