package service

import (
	_playerCoinModel "github.com/tehdev/summoner-rift-api/pkg/playerCoin/model"
)
type PlayerCoinService interface{
	CoinAdding(coinAddingReq *_playerCoinModel.CoinAddingReq) (*_playerCoinModel.PlayerCoin,error)
	Showing(playerID string)(*_playerCoinModel.PlayerCoinShowing, error)
}