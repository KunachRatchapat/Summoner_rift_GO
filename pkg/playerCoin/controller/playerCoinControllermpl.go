package controller

import (
	"net/http"
	"github.com/labstack/echo/v4"

	"github.com/tehdev/summoner-rift-api/pkg/custom"
	_playerCoinModel "github.com/tehdev/summoner-rift-api/pkg/playerCoin/model"
	_playerCoinService "github.com/tehdev/summoner-rift-api/pkg/playerCoin/service"
)

type playerCoinControllermpl struct{
		playerCoinService _playerCoinService.PlayerCoinService
}

func NewPlayerCoinControllermpl (playerCoinService _playerCoinService.PlayerCoinService) PlayerCoinController {
	return  &playerCoinControllermpl{playerCoinService: playerCoinService}
}

func (c *playerCoinControllermpl) CoinAdding(pctx echo.Context) error{
	coinAddingReq := new(_playerCoinModel.CoinAddingReq)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(coinAddingReq); err != nil{
		return  custom.Error(pctx, http.StatusBadRequest,err.Error())
	}
	
	playerCoin, err := c.playerCoinService.CoinAdding(coinAddingReq)
	if err != nil{
		return custom.Error(pctx, http.StatusInternalServerError, err.Error())
	}

	return  pctx.JSON(http.StatusCreated, playerCoin)
	

}